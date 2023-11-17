package utils

type utils struct{}

var _ Utils = &utils{}

func NewUtils() Utils {
	return &utils{}
}

func (*utils) Add(a, b int) int {
	return a + b
}

func (*utils) Sub(a, b int) int {
	panic("Implement me!")
}

func (*utils) Mul(a, b int) int {
	panic("Implement me!")
}

func (*utils) Div(a, b int) int {
	panic("Implement me!")
}
