package main

import "fmt"

type Turkish_Citizen struct {
	first, last, nationality string
}

type Russian_Citizen struct {
	first, last, nationality string
}

type American_Citizen struct {
	first, last, nationality string
}

func (t Turkish_Citizen) speak() {
	fmt.Printf("Ben %v %v ve Ben %v Vatandasiyim", t.first, t.last, t.nationality)
}

func (r Russian_Citizen) speak() {
	fmt.Printf("\nya %v %v ya grazhdanin %v", r.first, r.last, r.nationality)
}

func (a American_Citizen) speak() {
	fmt.Printf("\nI am %v %v and I am citizen of %v", a.first, a.last, a.nationality)
}

type Human interface {
	speak()
}

func nationality(h Human) {
	switch h.(type) {
	case Turkish_Citizen:
		fmt.Printf("\nBen %v %v ve Ben %v Vatandasiyim",
			h.(Turkish_Citizen).first, h.(Turkish_Citizen).last,
			h.(Turkish_Citizen).nationality)
	case Russian_Citizen:
		fmt.Printf("\nya %v %v ya grazhdanin %v",
			h.(Russian_Citizen).first, h.(Russian_Citizen).last,
			h.(Russian_Citizen).nationality)
	case American_Citizen:
		fmt.Printf("\nI am %v %v and I am Citizen of %v",
			h.(American_Citizen).first, h.(American_Citizen).last,
			h.(American_Citizen).nationality)
	}
}

func main() {
	t := Turkish_Citizen{
		first:       "Suzan",
		last:        "Kahramaner",
		nationality: "Turkiye Cumhuriyeti",
	}

	r := Russian_Citizen{
		first:       "Grigoriy",
		last:        "Perelman",
		nationality: "Rossiyskoy Federatsii",
	}

	a := American_Citizen{
		first:       "Katherine",
		last:        "Johnson",
		nationality: "United States of America",
	}

	t.speak()
	// Ben Suzan Kahramaner ve Ben Turkiye Cumhuriyeti vatandasiyim
	nationality(t)
	// Ben Suzan Kahramaner ve Ben Turkiye Cumhuriyeti vatandasiyim

	r.speak()
	// ya Grigoriy Perelman ya grazhdanin Rossiyskoy Federatsii
	nationality(r)
	// ya Grigoriy Perelman ya grazhdanin Rossiyskoy Federatsii

	a.speak()
	// I am Katherine Johnson and I'am citizen of United States of America
	nationality(a)
	// I am Katherine Johnson and I'am citizen of United States of America
}
