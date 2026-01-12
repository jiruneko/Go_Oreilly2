package main

import "fmt"

func main() {
	words := []string{"hi", "salutions", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "は短い単語です")
		case wordLen > 10:
			fmt.Println(word, "は長すぎる単語です")
		default:
			fmt.Println(word, "はちょうど良い長さの単語です")
		}
	}
}
