package main

import (
	"fmt"
)

type SchoolGirl struct {
	name string
}

type Proxy struct {
	gg Pursuit
}

func NewProxy() *Proxy {
	return &Proxy{Pursuit{"小菜", "娇娇"}}
}
func (p *Proxy) GiveDolls() {
	p.gg.GiveDolls()
	fmt.Println("戴利顺便放一张贺卡：祝你生日快乐！——戴利")
}

func (p *Proxy) GiveFlowers() {
	p.gg.GiveFlowers()
	fmt.Println("戴利顺便放一张贺卡：祝你生日快乐！——戴利")
}

func (p *Proxy) GiveChocolate() {
	p.gg.GiveChocolate()
	fmt.Println("戴利顺便放一张贺卡：祝你生日快乐！——戴利")
}

type GiveGift interface {
	giveDolls()
	giveFlowers()
	giveChocolate()
}

type Pursuit struct {
	name string
	mm   string
}

func (p *Pursuit) GiveDolls() {
	fmt.Println(p.name, "送", p.mm, "洋娃娃")
}

func (p *Pursuit) GiveFlowers() {
	fmt.Println(p.name, "送", p.mm, "鲜花")
}

func (p *Pursuit) GiveChocolate() {
	fmt.Println(p.name, "送", p.mm, "巧克力")
}

func main() {
	daili := NewProxy()
	daili.GiveDolls()
	daili.GiveFlowers()
	daili.GiveChocolate()
}