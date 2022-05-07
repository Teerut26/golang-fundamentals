package main

import (
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
		name: "holope",
		age:  18,
	}

	upperAllLetter(&p)
	increaseAge(&p, 5)

	println("name => ", p.name)
	println("age => ", p.age)
}

// ref https://dev.to/iporsut/go-pointer-pointer-go-3212
