package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	surname string
}

func (p person) getFullName() string {
	return p.name + " " + p.surname
}

func (p *person) setFullName(fullName string) {
	sections := strings.Split(fullName, " ")

	p.name = sections[0]
	p.surname = sections[1]
}

func main() {
	p1 := person{"John", "Doe"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Jane Doe")
	fmt.Println(p1.getFullName())
}
