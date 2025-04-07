package k8s

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
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

func getK8sClient() *kubernetes.Clientset {
	// 1. 获取kubeconfig路径
	kubeconfigPath := getKubeconfigPath()

	fmt.Println("kubeconfigPath:", kubeconfigPath)

	// 2. 创建配置
	config, err := createK8sConfig(kubeconfigPath, "")
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
