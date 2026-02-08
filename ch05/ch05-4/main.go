package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(num int, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return num, denom, errors.New("0で割ることはできません")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func main() {
	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)
}
