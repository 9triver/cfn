package models

import "github.com/9triver/cfn/cfn-go/messages"

type Resource struct {
	CPU    *CPU
	GPU    *GPU
	Memory *Memory
	Tags   []string
}

func (resource *Resource) toMessage() messages.Resource {
	return messages.Resource{
		CPU:    resource.CPU.toMessage(),
		GPU:    resource.GPU.toMessage(),
		Memory: resource.Memory.toMessage(),
		Tags:   resource.Tags,
	}
}

type CPU struct {
	Cores int32
}

func (cpu *CPU) toMessage() *messages.CPU {
	return &messages.CPU{
		Cores: cpu.Cores,
	}
}

type GPU struct {
	Cores int32
	Model string
}

func (gpu *GPU) toMessage() *messages.GPU {
	return &messages.GPU{
		Cores: gpu.Cores,
		Model: gpu.Model,
	}
}

type Memory struct {
	Capacity int32
}

const KB = 1024
const MB = 1024 * KB
const GB = 1024 * MB

func (memory *Memory) toMessage() *messages.Memory {
	return &messages.Memory{
		Capacity: memory.Capacity,
	}
}
