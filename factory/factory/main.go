package main

import "fmt"

func main()  {
	f := PlusOperatorFactory{}
	op := f.Create()
	op.SetLeft(10)
	op.SetRight(20)
	r := op.Result()
	fmt.Println(r)

	f1 := SubOperatorFactory{}
	op1 := f1.Create()
	op1.SetRight(20)
	op1.SetRight(10)
	r2 := op1.Result()
	fmt.Println(r2)

}
