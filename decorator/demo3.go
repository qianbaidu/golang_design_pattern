package main

import (
	"fmt"
)

type Finery interface {
	Show()
}

type Person struct {
	Name     string
	Fineries []Finery
}

func (this *Person) Show() {
	fmt.Println("装扮的", this.Name)
	for _, finery := range this.Fineries {
		finery.Show()
	}
}

func (this *Person) Decorate(finery Finery) {
	if this.Fineries == nil {
		this.Fineries = make([]Finery, 0, 0)
	}
	this.Fineries = append(this.Fineries, finery)
}

type TShirts struct {
	Finery
}

func (t *TShirts) Show() {
	fmt.Println("大T恤")
}

type BigTrouser struct {
	Finery
}

func (b *BigTrouser) Show() {
	fmt.Println("大裤衩")
}

type Sneakers struct {
	Finery
}

func (s *Sneakers) Show() {
	fmt.Println("破球鞋")
}

type LeatherShoes struct {
	Finery
}

func (l *LeatherShoes) Show() {
	fmt.Println("皮鞋")
}

type Suit struct {
	Finery
}

func (s *Suit) Show() {
	fmt.Println("西装")
}

type Tie struct {
	Finery
}

func (t *Tie) Show() {
	fmt.Println("领带")
}

func main() {
	person1 := &(Person{"小菜", nil})
	fmt.Println("第一种装扮：")
	dtx := new(TShirts)
	kk := new(BigTrouser)
	pqx := new(Sneakers)
	person1.Decorate(dtx)
	person1.Decorate(kk)
	person1.Decorate(pqx)
	person1.Show()

	person2 := &(Person{"小菜", nil})
	fmt.Println("第二种装扮：")
	px := new(LeatherShoes)
	ld := new(Tie)
	xz := new(Suit)
	person2.Decorate(xz)
	person2.Decorate(ld)
	person2.Decorate(px)
	person2.Show()

	person3 := &(Person{"小菜", nil})
	fmt.Println("第三种装扮：")
	xz2 := new(Suit)
	pqx2 := new(Sneakers)
	kk2 := new(BigTrouser)
	ld2 := new(Tie)
	person3.Decorate(xz2)
	person3.Decorate(ld2)
	person3.Decorate(kk2)
	person3.Decorate(pqx2)
	person3.Show()
}
