package com.github9triver.cfn.manager;

import com.github9triver.cfn.config.K8sClientProperties;
import com.github9triver.cfn.model.dto.ResourceDto;
import com.github9triver.cfn.model.dto.ServerAddress;
import com.github9triver.cfn.proto.data.Resources;
import io.kubernetes.client.custom.IntOrString;
import io.kubernetes.client.custom.Quantity;
import io.kubernetes.client.openapi.ApiClient;
import io.kubernetes.client.openapi.ApiException;
import io.kubernetes.client.openapi.apis.CoreV1Api;
import io.kubernetes.client.openapi.models.*;
import io.kubernetes.client.util.ClientBuilder;
import io.kubernetes.client.util.KubeConfig;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;

import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.math.BigDecimal;
import java.net.InetSocketAddress;
import java.net.Socket;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.function.Function;

@SuppressWarnings("SpringJavaInjectionPointsAutowiringInspection")
@Slf4j
public class K8sLocalResourceManager implements LocalResourceManager {

    private K8sClientProperties properties;

    @Autowired
    public void setProperties(K8sClientProperties properties) {
        this.properties = properties;
    }

    private volatile CoreV1Api api;

    private ApiClient createApiClient() {
        // 1. 加载 kubeconfig 文件
        KubeConfig kubeConfig;
        try {
            if (properties.getConfigFile() != null && !properties.getConfigFile().isEmpty()) {
                // 1.1 指定路径
                kubeConfig = KubeConfig.loadKubeConfig(new FileReader(properties.getConfigFile()));
            } else {
                // 1.2 默认路径
                String defaultKubeConfigPath = Paths.get(System.getProperty("user.home"), ".kube", "config").toString();
                kubeConfig = KubeConfig.loadKubeConfig(new FileReader(defaultKubeConfigPath));
            }
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        }
        // 2. 配置上下文
        kubeConfig.setContext(properties.getContext());
        try {
            // 3. 创建客户端
            return ClientBuilder.kubeconfig(kubeConfig).build();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public CoreV1Api getK8sApi() {
        if (api == null) {
            synchronized (this) {
                if (api == null) {
                    api = new CoreV1Api(createApiClient());
                }
            }
        }
        return api;
    }

    //    private final static String[] resourceTypes = {"cpu", "memory", "pods", "ephemeral-storage"};
//    private final static String[] resourceTypes = {"cpu", "memory"};

    private ResourceDto getResourceCount(ResourceAccessor accessor) {
        CoreV1Api k8sApi = getK8sApi();
        Map<String, BigDecimal> capacityMap = new HashMap<>();
        ResourceDto resourceDto = new ResourceDto();

        try {
            V1NodeList nodeList = k8sApi.listNode().execute();

            for (V1Node item : nodeList.getItems()) {
                Map<String, Quantity> capacity = accessor.apply(Objects.requireNonNull(item.getStatus()));
                if (capacity != null) {
                    for (Map.Entry<String, Quantity> entry : capacity.entrySet()) {
                        String key = entry.getKey();
                        BigDecimal value = entry.getValue().getNumber();
                        BigDecimal newValue = capacityMap.getOrDefault(key, BigDecimal.ZERO).add(value);
                        capacityMap.put(key, newValue);
                    }
                }
            }

        } catch (ApiException e) {
            throw new RuntimeException(e);
        }

        resourceDto.getCpu().setCores(capacityMap.get("cpu"));
        resourceDto.getMemory().setCapacity(capacityMap.get("memory"));

        return resourceDto;
    }

    @Override
    public ResourceDto getTotalResourceCount() {
        return getResourceCount(V1NodeStatus::getAllocatable);
    }

    @Override
    public ResourceDto getAvailableResourceCount() {
        Map<String, Map<String, BigDecimal>> nodeToAllocatable = new HashMap<>();
        Map<String, Map<String, BigDecimal>> nodeToUsed = new HashMap<>();

        // 获取所有 Node
        try {
            for (V1Node node : getK8sApi().listNode().execute().getItems()) {
                String nodeName = Objects.requireNonNull(node.getMetadata()).getName();
                Map<String, BigDecimal> alloc = new HashMap<>();
                for (Map.Entry<String, Quantity> entry : Objects.requireNonNull(node.getStatus()).getAllocatable().entrySet()) {
                    alloc.put(entry.getKey(), entry.getValue().getNumber());
                }
                nodeToAllocatable.put(nodeName, alloc);
                nodeToUsed.put(nodeName, new HashMap<>()); // 初始化使用记录
            }
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }

        // 获取所有 Pod
        try {
            for (V1Pod pod : getK8sApi().listPodForAllNamespaces().execute().getItems()) {
                String nodeName = Objects.requireNonNull(pod.getSpec()).getNodeName();
                if (nodeName == null || !nodeToUsed.containsKey(nodeName)) continue;

                for (V1Container container : pod.getSpec().getContainers()) {
                    Map<String, Quantity> requests = Objects.requireNonNull(container.getResources()).getRequests();
                    if (requests == null) continue;

                    Map<String, BigDecimal> used = nodeToUsed.get(nodeName);
                    for (Map.Entry<String, Quantity> entry : requests.entrySet()) {
                        String resType = entry.getKey();
                        BigDecimal existing = used.getOrDefault(resType, BigDecimal.ZERO);
                        used.put(resType, existing.add(entry.getValue().getNumber()));
                    }
                }
            }
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }

        // 输出结果
        ResourceDto resourceDto = new ResourceDto();

        for (String node : nodeToAllocatable.keySet()) {
            Map<String, BigDecimal> alloc = nodeToAllocatable.get(node);
            Map<String, BigDecimal> used = nodeToUsed.get(node);

            for (String key : ResourceDto.RESOURCE_TYPES) {
                BigDecimal allocQ = alloc.getOrDefault(key, BigDecimal.ZERO);
                BigDecimal usedQ = used.getOrDefault(key, BigDecimal.ZERO);
                BigDecimal remaining = allocQ.subtract(usedQ);
                resourceDto.set(key, remaining);
            }
        }

        return resourceDto;
    }


    interface ResourceAccessor extends Function<V1NodeStatus, Map<String, Quantity>> {
    }


    private final Object lock = new Object();
    private int podNumber = 0;

    /**
     * @return 唯一 pod 名
     */
    private String getUniqueName() {
        synchronized (lock) {
            return System.currentTimeMillis() + "-" + podNumber++;
        }
    }

    @Override
    public ServerAddress requestResources(Resources.Resource resource) {
        V1ResourceRequirements resources = new V1ResourceRequirements()
                .putRequestsItem("cpu", Quantity.fromString(resource.getCPU().getCores()))
                .putRequestsItem("memory", Quantity.fromString(resource.getMemory().getCapacity()))
                .putLimitsItem("cpu", Quantity.fromString(resource.getCPU().getCores()))
                .putLimitsItem("memory", Quantity.fromString(resource.getMemory().getCapacity()));

        String uniqueName = getUniqueName();
        String podName = "cfn-workenv-pod-" + uniqueName;
        String serviceName = "cfn-workenv-svc-" + uniqueName;
        Map<String, String> labels = Map.of(
                "app", "cfn-workenv",
                "name", podName
        );
        V1Pod pod = new V1Pod()
                .metadata(new V1ObjectMeta().name(podName).namespace("default").labels(labels))
                .spec(new V1PodSpec().overhead(null).containers(List.of(
                        new V1Container()
                                .name(getUniqueName())
                                .resources(resources)
                                .image("cfn-workenv-python:0.0")
                                .ports(List.of(new V1ContainerPort().containerPort(8667)))
                )));

        try {
            getK8sApi().createNamespacedPod("default", pod).execute();
        } catch (ApiException e) {
            log.error(e.toString());
            throw new RuntimeException(e);
        }

        V1Service service = new V1Service()
                .metadata(new V1ObjectMeta().name(serviceName).namespace("default"))
                .spec(new V1ServiceSpec()
                        .type("NodePort")
                        .selector(labels)  // 与 Pod 的 label 匹配
                        .ports(List.of(new V1ServicePort()
                                .port(8667)               // Service 的端口
                                .targetPort(new IntOrString(8667)) // Pod 容器的端口
                        ))
                );

        Integer assignedPort; // 获取由集群分配的访问端口
        try {
            V1Service createdService = api.createNamespacedService("default", service).execute();
            assignedPort = Objects.requireNonNull(
                    Objects.requireNonNull(
                            createdService.getSpec()
                    ).getPorts()
            ).getFirst().getNodePort();
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }

        if (assignedPort != null) {
            try {
                String address = getAccessibleAddress(assignedPort);

                return new ServerAddress(address, assignedPort);
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        } else {
            throw new RuntimeException("No assigned port");
        }

    }

    // 尝试连接 IP:port，看是否能访问
    @SuppressWarnings("SameParameterValue")
    private boolean isReachable(String ip, int port, int timeoutMillis) {
        try (Socket socket = new Socket()) {
            socket.connect(new InetSocketAddress(ip, port), timeoutMillis);
            return true;
        } catch (IOException e) {
            return false;
        }
    }

    // TODO: 缓存
    // 获取可访问的地址（优先 Node IP，不通则 fallback 到 host.docker.internal）
    private String getAccessibleAddress(int nodePort) throws Exception {
        V1NodeList nodeList = api.listNode().execute();

        for (V1Node node : nodeList.getItems()) {
            List<V1NodeAddress> addresses = Objects.requireNonNull(node.getStatus()).getAddresses();
            if (addresses != null) {
                for (V1NodeAddress addr : addresses) {
                    if ("ExternalIP".equals(addr.getType())) {
                        String ip = addr.getAddress();
                        if (isReachable(ip, nodePort, 1000)) {
                            return ip;
                        }
                    }
                }
            }
        }
        // fallback
        return "kubernetes.docker.internal";
    }
}