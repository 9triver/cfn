package task

type Function interface {
	Apply(map[string]any) (map[string]any, error)
}

type FunctionFactory func() Function
