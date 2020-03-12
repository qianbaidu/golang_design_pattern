package main

import (
	"fmt"
)

type OldInterface interface {
	InsertToDatabase(Data interface{}) (bool, error)
}

type AddCustomInfoToMysql struct {
	DbName string
}

func (pA *AddCustomInfoToMysql) InsertToDatabase(info interface{}) (bool, error) {
	switch info.(type) {
	case string:
		fmt.Println("add ", info.(string), " to ", pA.DbName, " successful!")
	}
	return true, nil
}

type NewInterface interface {
	SaveData(Data interface{}) (bool, error)
}

type Adapter struct {
	OldInterface
}

func (pA *Adapter) SaveData(Data interface{}) (bool, error) {
	fmt.Println("In Adapter")
	return pA.InsertToDatabase(Data)
}

func main() {
	// 老接口方法为InsertToDatabase，新接口方法为SaveData。两者具有一点的相似性。当系统不想保留接口的时候，就可以用适配器来修饰。
	var iNew NewInterface
	iNew = &Adapter{OldInterface: &AddCustomInfoToMysql{DbName: "mysql"}}
	iNew.SaveData("helloworld")

	// 旧接口调用，golang默认隐式的，所以只要实现了接口方法即可，未指定类型情况下实现了接口方法都可以调用
	o := &Adapter{&AddCustomInfoToMysql{"test"}}
	o.InsertToDatabase("old insert") //旧
	o.SaveData("new insert") //新
	o.OldInterface.InsertToDatabase("old ") //也可以显示调用旧方法

	return
}
