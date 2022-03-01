package main

import "fmt"

func main() {
	const constant_integer int = 10
	fmt.Println(constant_integer) // 10

	constant_integer = 20 // constant olarak degisken atamasi
	// yaptigimiz icin hataya sebep olur

	const constant_string string = "Hello"
	fmt.Println(constant_string) // Hello

	constant_string = "World" // constant olarak degisken atamasi
	// yaptigimiz icin hataya sebep olur
}
