package singleton

import (
	"fmt"
	"sync"
)

type Single struct {
	data int
}

var singleton *Single
var once sync.Once

func GetInterface() *Single {
	//once.Do(func() {
	//	singleton = &Single{100}
	//})
	singleton = &Single{100}
	return singleton
}

func main() {
	i1 := GetInterface()
	i2 := GetInterface()
	if i1 == i2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
}
