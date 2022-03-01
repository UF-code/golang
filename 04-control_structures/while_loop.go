package main

import "fmt"

func main() {
	x := 0
	for x <= 5 {
		fmt.Printf("%d ", x) // 0 1 2 3 4 5
		x++
	}
}
