package main

func main() {

	e := NewCommonMessage(ViaEmail())
	e.SendMessage("test ", "xxx@xxx.com")

	e1 := NewCommonMessage(ViaSms())
	e1.SendMessage("花费查询", "10086")

	e2 := NewErrorMessage(ViaSms())
	e2.SendMessage("system error ","138xxxx")
}
