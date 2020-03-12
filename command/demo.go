package main

import "fmt"

/**
Receiver角色实现
*/
type TV struct{}

func (tv TV) Open() {
	fmt.Println("开机")
}

func (tv TV) Close() {
	fmt.Println("关机")
}

func (tv TV) ChangeChannel() {
	fmt.Println("换台")
}

/**
Command以及ConcreteCommand角色实现
*/

type Command interface {
	Execute()
}

type OpenCommand struct {
	receiver TV
}

func (oc OpenCommand) Execute() {
	oc.receiver.Open()
}

type CloseCommand struct {
	receiver TV
}

func (cc CloseCommand) Execute() {
	cc.receiver.Close()
}

type ChangeChannelCommand struct {
	receiver TV
}

func (ccc ChangeChannelCommand) Execute() {
	ccc.receiver.ChangeChannel()
}

/**
Invoker角色实现
*/

type Invoke struct {
	Command
}

func (i Invoke) ExecuteCommand() {
	i.Command.Execute()
}

/**
定义工厂方法创建Command
*/
func NewCommand(t string, tv TV) Command {
	switch t {
	case "open":
		return OpenCommand{
			receiver: tv,
		}
	case "close":
		return CloseCommand{
			receiver: tv,
		}
	case "changechannel":
		return ChangeChannelCommand{
			receiver: tv,
		}
	default:
		return nil
	}
}

/**
客户端调用
*/
func main() {
	//创建一个Reveiver
	tTV := TV{}
	//创建一个Command
	tCommand := NewCommand("open", tTV)
	//创建一个调用者
	tInvoke := Invoke{
		Command: tCommand,
	}
	tInvoke.ExecuteCommand()


	tCommand = NewCommand("changechannel", tTV)
	tInvoke = Invoke{
		Command: tCommand,
	}
	tInvoke.ExecuteCommand()


	tCommand = NewCommand("close", tTV)
	tInvoke = Invoke{
		Command: tCommand,
	}
	tInvoke.ExecuteCommand()
}
