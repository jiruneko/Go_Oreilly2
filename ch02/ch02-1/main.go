package main

import (
	"fmt"
)

func main() {
	var flag bool
	var isAwesome = true
	fmt.Println(isAwesome)
	fmt.Println(flag)
	var complexNum = complex(20.3, 10.2)
	println(complexNum)
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
}
