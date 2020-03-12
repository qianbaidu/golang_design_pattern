package main

import "fmt"

type ComputerInput interface {
	Input()
}
type Keyboard struct {
}

func (Keyboard) Input() {
	fmt.Println("usb input ")
}

type Mouse struct {
}

func (Mouse) Input() {
	fmt.Println("mouse input")
}

// 由一个工厂对象决定创建出哪一种产品类的实例。定义一个创建对象的类，由这个类来封装实例化对象的行为
func SimpleFactory(str string) ComputerInput {
	switch str {
	case "usb":
		return &Keyboard{}
	default:
		return &Mouse{}
	}
}

func main() {
	SimpleFactory("usb").Input()
}
