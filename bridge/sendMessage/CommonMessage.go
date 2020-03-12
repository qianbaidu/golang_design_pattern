package main

type CommonMessage struct {
	method MessageImp
}

func NewCommonMessage(method MessageImp) *CommonMessage {
	return &CommonMessage{method: method}
}

func (com *CommonMessage) SendMessage(text, to string) {
	com.method.send(text, to)
}

