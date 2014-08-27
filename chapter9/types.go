package main

import "fmt"

type Person struct {
	Name string
}

type Speakable interface {
	Talk()
}


func (p *Person) Talk() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p *Android) Talk() {
	fmt.Println("*brrr* Hello, my name is", p.Name)
}


func chat(chatters ...Speakable) {
	for _, v := range chatters {
		v.Talk()
	}
}

type Android struct {
	Person
	Model string
}

func main() {
	p := Person{"Jabbslad"}
	a := Android{Person: Person{"Jabbsdroid"}, Model: "GS11"}

	chat(&p, &a)
}