// 模板方法模式
package main

import "fmt"

type Template interface {
	fun1()
	fun2()
}

//此处采取过程式编程，因为Go没有抽象类的概念
func result(t Template) {
	t.fun1()
	t.fun2()
}

type ConcreteA struct{}

func (c *ConcreteA) fun1() {
	fmt.Println("A类实现fun1")
}

func (c *ConcreteA) fun2() {
	fmt.Println("A类实现fun2")
}

type ConcreteB struct{}

func (c *ConcreteB) fun1() {
	fmt.Println("B类实现fun1")
}

func (c *ConcreteB) fun2() {
	fmt.Println("B类实现fun2")
}

func main() {
	// var ta = new(ConcreteA)
	// result(ta)

	var tb = new(ConcreteB)
	result(tb)
}