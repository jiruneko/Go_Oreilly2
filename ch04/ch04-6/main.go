package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("fizz buzz")
			} else {
				fmt.Println("fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
