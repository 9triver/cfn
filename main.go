package main

import (
	"fmt"
	"github.com/kekwy/cfn/node"
)

func main() {
	fmt.Println("Hello World")
	pid1 := node.RunHeadNode("localhost", 8080, "headnode", []string{})
	pid2 := node.RunHeadNode("localhost", 8081, "headnode", []string{"headnode@127.0.0.1:8080"})
	fmt.Println(pid1)
	fmt.Println(pid2)
	//pid3 := actor.NewPID("localhost:8080", "$3")

	select {}
}
