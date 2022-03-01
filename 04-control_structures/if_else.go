package main

import "fmt"

func main() {
	var num int = 10
	var num2 int = 20

	if num == num2 {
		fmt.Printf("%d == %d \n", num, num2)
	} else if num != num2 {
		fmt.Printf("%d != %d \n", num, num2)
	} else {
		fmt.Println("Eger hic biri dogru degilse bu satir calisir")
	}
}
