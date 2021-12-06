package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name     string
	lastname string
}

func (p Person) getFullName() string {
	return p.name + " " + p.lastname
}

func (p *Person) setFullName(fullName string) {
	splits := strings.Split(fullName, " ")
	p.lastname = splits[len(splits)-1]
	p.name = strings.Join(splits[:len(splits)-1], " ")
}

func main() {
	person := Person{}

	person.setFullName("Andre Vitor Rampanelli")

	fmt.Println(person.getFullName())
}
