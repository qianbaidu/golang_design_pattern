package main

import "fmt"

type Api interface {
	Say(name string) string
}

func NewApi(str string) Api  {
	switch str {
	case "english":
		return &English{}
	default:
		return &Chinese{}
	}

}

type Chinese struct {
}

func (*Chinese) Say(name string) string {
	return "您好 " + name
}

type English struct {

}

func (*English) Say(name string) string {
	return "hello " + name
}

func main() {
	api := NewApi("chinese")
	res := api.Say("张三")
	fmt.Println(res)

	api = NewApi("english")
	res = api.Say("LiLei")
	fmt.Println(res)
}
