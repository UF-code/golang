package main

import "fmt"

func change_the_value_of_pointer(ptr *int, value int) {
	*ptr = value
}

func main() {
	value_1 := 10
	fmt.Println(value_1)
	// 10
	fmt.Println(&value_1)
	// 0xc0000aa058

	value_2 := 20
	change_the_value_of_pointer(&value_1, value_2)
	fmt.Println(value_1)
	// 20
}
