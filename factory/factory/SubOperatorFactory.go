package main

type SubOperatorFactory struct {
}
type SubOperator struct {
	*OperatorBase
}

func (o SubOperator) Result() int {
	return o.left - o.right
}
func (SubOperatorFactory) Create() Operator {
	return SubOperator{&OperatorBase{}}
}
