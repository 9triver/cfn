package k8s

import (
	"context"
	"fmt"
	messages "github.com/9triver/ignis/proto/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getKubeconfigPath() string {
	// 优先级1：环境变量指定的路径
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		return kubeconfig
	}

	// 优先级2：默认的 kubeconfig 路径
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Warning: Failed to get home directory: %v", err)
		return ""
	}
	return filepath.Join(home, ".kube", "config")
}

func createK8sConfig(kubeconfigPath, contextName string) (*rest.Config, error) {
	// 如果 kubeconfigPath 为空，尝试使用 in-cluster 配置
	if kubeconfigPath == "" {
		log.Println("No kubeconfig path specified, trying in-cluster configuration")
		return rest.InClusterConfig()
	}

	// 使用指定的 kubeconfig 和上下文
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = kubeconfigPath

	configOverrides := &clientcmd.ConfigOverrides{}
	if contextName != "" {
		configOverrides.CurrentContext = contextName
	}

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides,
	)

	return kubeConfig.ClientConfig()
}

const ContextName = "docker-desktop"

func getK8sClient() *kubernetes.Clientset {
	// 1. 获取kubeconfig路径
	kubeconfigPath := getKubeconfigPath()

	fmt.Println("kubeconfigPath:", kubeconfigPath)

	// 2. 创建配置
	config, err := createK8sConfig(kubeconfigPath, ContextName)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes config: %v", err)
	}

	// 3. 创建客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	return clientset
}

func getK8sNode(clientset *kubernetes.Clientset, nodeName string) *corev1.Node {
	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return node
}

const (
	PodName  = "cfn-work-app"
	PodImage = "cfn-work-platform-python:0.0.0"
	PodPort  = 8667
)

func CreatePodAndService(requestResource *messages.Resource) (string, int32, error) {
	clientset := getK8sClient()
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      PodName,
			Namespace: "default",
			Labels: map[string]string{
				"app": PodName,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  PodName,
					Image: PodImage,
					Ports: []corev1.ContainerPort{
						{
							Name:          "platform-actor",
							ContainerPort: PodPort,
							Protocol:      corev1.ProtocolTCP,
						},
					},
				},
			},
		},
	}

	pod.Spec.Containers[0].Resources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			"cpu":    resource.MustParse(requestResource.GetCPU().GetCores()),
			"memory": resource.MustParse(requestResource.GetMemory().GetCapacity()),
		},
		Limits: corev1.ResourceList{
			"cpu":    resource.MustParse(requestResource.GetCPU().GetCores()),
			"memory": resource.MustParse(requestResource.GetMemory().GetCapacity()),
		},
	}

	// 4. 创建 Pod
	_, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create pod: %v", err))
	}
	fmt.Printf("Pod %s created successfully\n", PodName)

	// 5. 创建 Service 暴露 Pod
	serviceName := fmt.Sprintf("%s-svc", PodName)
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: "default",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": PodName, // 匹配 Pod 的标签
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "platform-actor",
					Port:       PodPort,
					TargetPort: intstr.FromInt32(PodPort),
					Protocol:   corev1.ProtocolTCP,
					// 不指定 NodePort 让 Kubernetes 自动分配
				},
			},
		},
	}

	// 6. 创建 Service
	createdService, err := clientset.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create service: %v", err))
	}

	// 7. 获取分配的 NodePort
	nodePort := createdService.Spec.Ports[0].NodePort
	fmt.Printf("Service created with NodePort: %d\n", nodePort)

	// 8. 获取节点 IP
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		panic("Failed to get node information")
	}

	nodeIP := ""
	for _, addr := range nodes.Items[0].Status.Addresses {
		if addr.Type == corev1.NodeExternalIP {
			nodeIP = addr.Address
			break
		}
		if addr.Type == corev1.NodeInternalIP {
			nodeIP = addr.Address
		}
	}

	if strings.Contains(nodes.Items[0].Name, "docker-desktop") {
		nodeIP = "127.0.0.1"
	}

	if nodeIP == "" {
		panic("No valid node IP found")
	}

	// 9. 输出访问信息
	fmt.Printf("\nPod access information:\n")
	fmt.Printf("Pod Name: %s\n", PodName)
	fmt.Printf("Access URL: http://%s:%d\n", nodeIP, nodePort)
	fmt.Printf("To delete: kubectl delete pod %s -n default\n", PodName)
	fmt.Printf("To delete service: kubectl delete svc %s -n default\n", serviceName)

	return nodeIP, nodePort, nil
}

func getNodeMetrics(clientset *kubernetes.Clientset, nodeName string) {
	// 注意: 需要使用 metrics.k8s.io/v1beta1 API
	// 需要额外导入:
	// "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	// metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"

	// 创建metrics客户端
	//metricsConfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	////metricsClient, err := metricsclientset.NewForConfig(metricsConfig)
	//if err != nil {
	//	panic(err.Error())
}

////nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().Get(context.TODO(), nodeName, metav1.GetOptions{})
////if err != nil {
////	panic(err.Error())
////}
//
//fmt.Printf("\nNode Metrics for %s:\n", nodeName)
//fmt.Printf("CPU Usage: %v\n", nodeMetrics.Usage.Cpu().String())
//fmt.Printf("Memory Usage: %v\n", nodeMetrics.Usage.Memory().String())
//}
