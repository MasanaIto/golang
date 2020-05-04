package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = p.Name
	fmt.Println(p.Name)
	return p.Name
}

func Access(human Human) {
	if human.Say() == "Masana" {
		fmt.Println("OK, WELCOME MASANA.")
	} else {
		fmt.Println("You do not have access")
	}
}

func main() {
	var masana Human = &Person{"Masana"}
	masana.Say() // Mr.Masana
	Access(masana)
}
