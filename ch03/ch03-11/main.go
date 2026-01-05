package main

func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "ボブ"
	person.age = 50
	person.pet = "dog"

	// pet := struct {
	// 	name string
	// 	kind string
	// }{
	// 	name: "ポチ",
	// 	kind: "dog",
	// }
}
