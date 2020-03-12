package main

import "fmt"

type EmailMessage struct {
}

func ViaEmail() MessageImp {
	return &EmailMessage{}
}

func (m *EmailMessage) send(text, to string) {
	fmt.Println(fmt.Sprintf("send email :[ %s ] to %s ", text, to))
}
