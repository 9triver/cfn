package main

import (
	"fmt"
	"github.com/kekwy/cfn/k8s"
	"log"
	"os"
	"path/filepath"
)

func GetKubeconfigPath() string {
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

func main() {
	fmt.Println("Hello World")
	manager := k8s.NewLocalResourceManager()
	manager.CreateWorkerActor(nil)
	//pid1 := node.RunHeadNode("localhost", 8080, "headnode", []string{})
	//pid2 := node.RunHeadNode("localhost", 8081, "headnode", []string{"headnode@127.0.0.1:8080"})
	//fmt.Println(pid1)
	//fmt.Println(pid2)
	//pid3 := client.NewPID("localhost:8080", "$3")

	//fmt.Println(GetKubeconfigPath())

	//select {}
}
