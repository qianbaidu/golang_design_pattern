package main

import "fmt"

/**
实现Mediator以及ConcreteMediator类
*/

type UnitedNations interface {
	ForwardMessage(message string, country Country)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Iraq
}

func (unsc UnitedNationsSecurityCouncil) ForwardMessage(message string, country Country) {
	switch country.(type) {
	case USA:
		unsc.Iraq.GetMessage(message)
	case Iraq:
		unsc.USA.GetMessage(message)
	default:
		fmt.Println("The country is not a member of UNSC")
	}

}

/**
实现Colleague以及ConcreteColleague类
 */
type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}
type USA struct {
	UnitedNations
}

func (usa USA) SendMessage(message string) {
	usa.UnitedNations.ForwardMessage(message, usa)
}

func (usa USA) GetMessage(message string) {
	fmt.Printf("美国收到对方消息: %s\n", message)
}

type Iraq struct {
	UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
	iraq.UnitedNations.ForwardMessage(message, iraq)
}

func (iraq Iraq) GetMessage(message string) {
	fmt.Printf("伊拉克收到对方消息: %s\n", message)
}

func main() {
	//创建一个具体中介者
	tMediator := UnitedNationsSecurityCouncil{}
	//创建具体同事,并且让他认识中介者
	tColleageA := USA{
		UnitedNations: tMediator,
	}
	tColleageB := Iraq{
		UnitedNations: tMediator,
	}
	//让中介者认识每一个具体同事
	tMediator.USA = tColleageA
	tMediator.Iraq = tColleageB
	//A同事发送消息
	tColleageA.SendMessage("停止核武器研发，否则发动战争")
	tColleageB.SendMessage("我们没有研发核武器，也不怕战争")
}
