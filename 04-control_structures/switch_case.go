package main

import "fmt"

func main() {
	switch "Carsamba" {
	case "Pazartesi":
		fmt.Println("Gunlerden Pazartesi")
	case "Sali":
		fmt.Println("Gunlerden Sali")
	case "Carsamba":
		fmt.Println("Gunlerden Carsamba") // Gunlerden Carsamba
	case "Persembe":
		fmt.Println("Gunlerden Persembe")
	case "Cuma":
		fmt.Println("Gunlerden Cuma")
	case "Cumartesi":
		fmt.Println("Gunlerden Cumartesi")
	default:
		fmt.Println("Gunlerden Pazar")
	}
}
