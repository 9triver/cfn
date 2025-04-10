package main

import (
	"fmt"
	"github.com/kekwy/cfn/task"
)

type AddFunction struct {
	sum int
}

func (a *AddFunction) Apply(params map[string]any) (map[string]any, error) {
	fmt.Println(params)
	nums := params["nums"].([]int)

	for _, num := range nums {
		a.sum += num
	}
	return map[string]any{"sum": a.sum}, nil
}

func GetFunction() task.Function {
	return &AddFunction{
		sum: 0,
	}
}
