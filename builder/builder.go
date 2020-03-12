package main

// 抽象建造者
type Builder interface {
	Step1()
	Step2()
	Step3()
}

type Director struct {
	builder Builder
}


func NewDirector(b Builder) *Director {
	return &Director{b}
}

func (d *Director) Build() {
	d.builder.Step1()
	d.builder.Step2()
	d.builder.Step3()
}
