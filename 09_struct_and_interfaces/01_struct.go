package main

import "fmt"

type person struct {
	first      string
	last       string
	age        int
	founder_of string
}

func main() {
	person_1 := person{
		first:      "Satoshi",
		last:       "Nakamoto",
		age:        32,
		founder_of: "Bitcoin",
	}

	fmt.Println(person_1)
	// {Satoshi Nakomoto 32 BTC}
	first_name_1 := person_1.first + " " + person_1.last
	fmt.Printf("Full Name: %v Age: %v Founder of: %v \n",
		first_name_1, person_1.age, person_1.founder_of)
	// Full Name: Satoshi Nakomoto Age: 32 Founder of: BTC

	person_2 := person{
		first:      "Vitalik",
		last:       "Buterin",
		age:        28,
		founder_of: "Ethereum",
	}

	fmt.Println(person_2)
	// {Vitalik Buterin 28 ETH}
	first_name_2 := person_2.first + " " + person_2.last
	fmt.Printf("Full Name: %v Age: %v Founder of: %v \n",
		first_name_2, person_2.age, person_2.founder_of)
	// Full Name: Vitalik Buterin Age: 28 Founder of: ETH
}
