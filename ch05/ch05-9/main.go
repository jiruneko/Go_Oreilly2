package main

import "fmt"

var (
	add = func(i int, j int) int { return i + j }
	sub = func(i int, j int) int { return i - j }
	mul = func(i int, j int) int { return i * j }
	div = func(i int, j int) int { return i / j }
)

func main() {
	x := add(3, 4)
	fmt.Println(x)
	changeAdd()
	y := add(4, 6)
	fmt.Println(y)
}

func changeAdd() {
	add = func(i int, j int) int { return i + j + j }
}
