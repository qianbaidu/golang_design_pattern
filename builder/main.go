package main

func main() {
	s := &StringBuilder{}
	d := NewDirector(s)
	d.Build()

	i := &IntBuilder{}
	d := NewDirector(i)
	d.Build()

}
