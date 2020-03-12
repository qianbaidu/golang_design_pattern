package main

import (
	"fmt"
)

// 享元对象接口
type IFlyweight interface {
	Operation(int) //来自外部的状态
}

// 共享对象
type ConcreteFlyweight struct {
	name string
}

func (c *ConcreteFlyweight) Operation(outState int) {
	if c == nil {
		return
	}
	fmt.Println("共享对象响应外部状态", outState)
}

// 不共享对象
type UnsharedConcreteFlyweight struct {
	name string
}

func (c *UnsharedConcreteFlyweight) Operation(outState int) {
	if c == nil {
		return
	}
	fmt.Println("不共享对象响应外部状态", outState)
}

// 享元工厂对象
type FlyweightFactory struct {
	flyweights map[string]IFlyweight
}

func (f *FlyweightFactory) Flyweight(name string) IFlyweight {
	if f == nil {
		return nil
	}
	if name == "u" {
		return &UnsharedConcreteFlyweight{"u"}
	} else if _, ok := f.flyweights[name]; !ok {
		f.flyweights[name] = &ConcreteFlyweight{name}
	}
	return f.flyweights[name]
}

func NewFlyweightFactory() *FlyweightFactory {
	ff := FlyweightFactory{make(map[string]IFlyweight)}
	ff.flyweights["a"] = &ConcreteFlyweight{"a"}
	ff.flyweights["b"] = &ConcreteFlyweight{"b"}
	return &ff
}

func main()  {
	ff := NewFlyweightFactory()

	fya := ff.Flyweight("a")
	fya.Operation(1)

	fyb := ff.Flyweight("b")
	fyb.Operation(2)

	fyc := ff.Flyweight("c")
	fyc.Operation(3)

	fyd := ff.Flyweight("d")
	fyd.Operation(4)

	fyu := ff.Flyweight("u")
	fyu.Operation(5)
}