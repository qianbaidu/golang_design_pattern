package main

// 抽象
type Message interface {
	SendMessage(text, to string)
}

// 实现
type MessageImp interface {
	send(text, to string)
}

// 通过将抽象和实现分开，可以分别不同扩展和修改