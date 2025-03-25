package main

import (
	"fmt"
	"github.com/kekwy/cfn/core/node"
)

func main() {
	fmt.Println("Hello World")
	pid1 := node.RunCommonHeadNode("localhost", 8080, "headnode", []string{})
	pid2 := node.RunCommonHeadNode("localhost", 8081, "headnode", []string{"headnode@localhost:8080"})
	fmt.Println(pid1)
	fmt.Println(pid2)
	//pid3 := actor.NewPID("localhost:8080", "$3")

	select {}
}
