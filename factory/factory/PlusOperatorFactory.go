package main

//操作抽象
type PlusOperatorFactory struct {
}
type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.left + o.right
}

func (PlusOperatorFactory) Create() Operator {
	return PlusOperator{&OperatorBase{}}
}


