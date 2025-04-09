package main

import (
	"fmt"
	"github.com/kekwy/cfn/client"
	"google.golang.org/protobuf/types/known/structpb"
)

type AddFunction struct {
	sum int
}

func (a *AddFunction) Apply(params *client.TaskInputs) (*client.TaskOutputs, error) {
	fmt.Println(params)
	nums := params.Fields["nums"].GetListValue()

	for _, num := range nums.Values {
		a.sum += int(num.GetNumberValue())
	}

	res, err := structpb.NewStruct(map[string]interface{}{
		"sum": a.sum,
	})
	if err != nil {
		return nil, err
	}
	return (*client.TaskOutputs)(res), nil
}

func main() {
	env := client.GetEnvironment()
	env.DeployFunction(&AddFunction{
		sum: 0,
	})
	env.Execute()
}
