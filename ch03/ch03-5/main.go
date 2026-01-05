package main

import (
	"fmt"
)

func main() {
	s := []string{"first", "socond", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))
	fmt.Printf("s[0]=|%s|, s[1]=|%s|, s[2]=|%s|\n", s[0], s[1], s[2])
}
