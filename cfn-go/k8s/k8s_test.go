package k8s

import (
	rs "github.com/9triver/ignis/proto/resource"
	"testing"
)

func TestCreatePodAndService(t *testing.T) {
	CreatePodAndService(&rs.Resource{
		CPU: &rs.CPU{
			Cores: "2000m",
		},
		Memory: &rs.Memory{
			Capacity: "1024Mi",
		},
		Tags: nil,
	})
}
