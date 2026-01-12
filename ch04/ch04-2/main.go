package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		a, x := 5, 20
		fmt.Println(a, x)
	}
	fmt.Println(x)
}
