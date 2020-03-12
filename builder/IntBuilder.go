package main

import "fmt"

type IntBuilder struct {
	result int64
}

func (s *IntBuilder) Step1() {
	s.result += 1
	fmt.Println(s.result)
}

func (s *IntBuilder) Step2() {
	s.result += 2
	fmt.Println(s.result)
}

func (s *IntBuilder) Step3() {
	s.result += 3
	fmt.Println(s.result)
}

func (s *IntBuilder) Build() string {
	return fmt.Sprintf("result:  %d", s.result)
}
