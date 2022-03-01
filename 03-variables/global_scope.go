package main

import "fmt"

var global_integer int = 10
var global_string string = "Hello World"
var global_bool bool = true

func main() {
	fmt.Println(global_integer) // 10
	fmt.Println(global_string)  // Hello World
	fmt.Println(global_bool)    // true
}
