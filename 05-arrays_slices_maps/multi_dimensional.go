package main

import "fmt"

func main() {
	first_matrix := []string{"Hello", "World", "With", "Go"}
	fmt.Println(first_matrix)

	second_matrix := []string{"Golang", "Coding", "is", "Fun"}
	fmt.Println(second_matrix)

	multi_dimensional := [][]string{first_matrix, second_matrix}
	fmt.Println(multi_dimensional) // [[Hello World With Go] [Golang Coding is Fun]]

}
