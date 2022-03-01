package main

import "fmt"

func main() {
	int_slice := []int{1, 2, 3, 4, 5}
	fmt.Println(int_slice) // [1 2 3 4 5]

	int_slice = append(int_slice, 6, 7, 8, 9)
	fmt.Println(int_slice) // [1 2 3 4 5 6 7 8 9]

	str_slice := []string{"Hello", "World", "With", "Golang"}
	fmt.Println(str_slice) // [Hello World With Golang]

	str_slice = append(str_slice, "Appending", "New", "Words")
	fmt.Println(str_slice) // [Hello World With Golang Appending New Words]

}
