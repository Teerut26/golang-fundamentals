package main

import (
	"fmt"
)

func main() {
	deno := []int{1, 2, 5, 10, 20, 50, 100, 500, 1000}

	ans := []int{}

	value := 0

	fmt.Print("value : ")
	fmt.Scan(&value)

	i := len(deno) - 1

	for i >= 0 {
		for value >= deno[i] {
			value -= deno[i]
			ans = append(ans, deno[i])
		}
		i -= 1
	}

	fmt.Println("-------แบงค์ที่ต้องทอน-------")
	for i, s := range ans {
		fmt.Println(i+1, ". ", s, " บาท")
	}

}
