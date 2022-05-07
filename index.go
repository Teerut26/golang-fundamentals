package main

import (
	"fmt"
	"time"
)

func calulate(x int, y int) int {
	return x + y
}

func showArray(list_name []string) {
	for i := 0; i < len(list_name); i++ {
		fmt.Println("[", i, "] => ", list_name[i])
	}
}

func showHelloWorld(n int, word string, delay int) {
	for i := 0; i < n; i++ {
		fmt.Println(word, " => ", i+1)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func showPiramid(layer int) {
	for i := 0; i < layer; i++ {
		if i != 0 {
			for j := layer - i; j > 1; j-- {
				fmt.Print(" ")
			}
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func main() {
	// n := 0
	// word := "Hello World"

	// list_name := []string{"b1", "b2"}

	// fmt.Print("hello word lines : ")
	// fmt.Scan(&n)

	// showHelloWorld(n, word, 1)

	// showArray(list_name)

	showPiramid(10)

}
