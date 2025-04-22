package com.github9triver.cfn.manager;

import com.github9triver.cfn.config.K8sClientProperties;
import io.kubernetes.client.openapi.ApiClient;
import io.kubernetes.client.openapi.ApiException;
import io.kubernetes.client.openapi.apis.CoreV1Api;
import io.kubernetes.client.openapi.models.V1Node;
import io.kubernetes.client.openapi.models.V1NodeList;
import io.kubernetes.client.util.ClientBuilder;
import io.kubernetes.client.util.KubeConfig;
import org.springframework.beans.factory.annotation.Autowired;

import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Paths;
import java.util.Objects;

@SuppressWarnings("SpringJavaInjectionPointsAutowiringInspection")
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
    private final static String[] resourceTypes = {"cpu", "memory"};

    public void getAllResources() {
        CoreV1Api k8sApi = getK8sApi();
        try {
            V1NodeList nodeList = k8sApi.listNode().execute();
            for (V1Node item : nodeList.getItems()) {
                System.out.println(Objects.requireNonNull(item.getStatus()).getCapacity());
            }
        } catch (ApiException e) {
            throw new RuntimeException(e);
        }
    }

}
