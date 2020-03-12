package main

import "fmt"

//实际运行类的接口
type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result() int
}

//工厂接口
type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	left, right int
}

//赋值
func (op *OperatorBase) SetLeft(left int) {
	op.left = left
}

func (op *OperatorBase) SetRight(right int) {
	op.right = right
}

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

type SubOperatorFactory struct {
}
type SubOperator struct {
	*OperatorBase
}

func (o SubOperator) Result()int  {
	return o.left - o.right
}
func (SubOperatorFactory) Create() Operator {
	return SubOperator{&OperatorBase{}}
}


func main() {
	f := PlusOperatorFactory{}
	op := f.Create()
	op.SetLeft(10)
	op.SetRight(20)
	r := op.Result()
	fmt.Println(r)

	f1 := SubOperatorFactory{}
	op1 := f1.Create()
	op1.SetRight(20)
	op1.SetRight(10)
	r2 := op1.Result()
	fmt.Println(r2)

}
