package main

import "fmt"

type ErrorMessage struct {
	method MessageImp
}

func NewErrorMessage(m MessageImp) *ErrorMessage {
	return &ErrorMessage{m}
}

func (e ErrorMessage) SendMessage(text, to string) {
	e.method.send(fmt.Sprintf("error ： %s", text), to)
}
