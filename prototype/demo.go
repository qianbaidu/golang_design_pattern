package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

//原型对象的类
type ProtoTypeManger struct {
	protoTypes map[string]Cloneable
}

func NewProtoTypeManger() *ProtoTypeManger {
	return &ProtoTypeManger{make(map[string]Cloneable)}
}

func (p *ProtoTypeManger) Get(name string) Cloneable {
	return p.protoTypes[name]
}

func (p *ProtoTypeManger) Set(name string, protoType Cloneable) {
	p.protoTypes[name] = protoType
}

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc //深复制
	//return t
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t //开辟内存新建变量，存储指针的内荣
	return &tc
}

func main() {
	mgr := NewProtoTypeManger()
	t1 := Type1{name: "type1"}
	mgr.Set("t1", &t1)
	t11 := mgr.Get("t1")
	t22 := t11.Clone()
	if t11 == t22 {
		fmt.Println("浅复制")
	} else {
		fmt.Println("深复制")
	}

}
