package main

import "fmt"

type StringBuilder struct {
	result string
}

func (s *StringBuilder) Step1() {
	s.result += "step 1 "
	fmt.Println(s.result)
}

func (s *StringBuilder) Step2() {
	s.result += "step 2 "
	fmt.Println(s.result)
}

func (s *StringBuilder) Step3() {
	s.result += "step 3 "
	fmt.Println(s.result)
}

//func (s *StringBuilder) Build() string {
//	return "result: " + s.result
//}
