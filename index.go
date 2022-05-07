package main

import (
	"fmt"
	"strings"
)

type Profile struct {
	name string
	age  int
}

func upperAllLetter(p *Profile) {
	*&p.name = strings.ToUpper(*&p.name)
}

func increaseAge(p *Profile, increaseValue int) {
	*&p.age = p.age + increaseValue
}

func main() {
	p := Profile{
		name: "holilope",
		age:  18,
	}

	upperAllLetter(&p)
	increaseAge(&p, 5)

	fmt.Println("name => ", p.name, " => ", &p.name)
	fmt.Println("age => ", p.age, " => ", &p.age)
}

// ref https://dev.to/iporsut/go-pointer-pointer-go-3212
