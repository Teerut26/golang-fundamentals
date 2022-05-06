package main

import (
	"fmt"
)

func main() {
	n := 0
	word := "Hello World"

	list_name := [...]string{"b1", "b2"}

	fmt.Print("hello word lines : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println(word, " => ", i+1)
	}

	for i := 0; i < len(list_name); i++ {
		fmt.Println("[", i, "] => ", list_name[i])
	}

}
