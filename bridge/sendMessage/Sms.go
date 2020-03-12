package main

import "fmt"

type SmsMessage struct {
}

func ViaSms() MessageImp {
	return &SmsMessage{}
}

func (v *SmsMessage) send(text, to string) {
	fmt.Println(fmt.Sprintf("send sms :[ %s ] to %s ", text, to))
}
