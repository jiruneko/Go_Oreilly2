package main

import "fmt"

func main() {
	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}
