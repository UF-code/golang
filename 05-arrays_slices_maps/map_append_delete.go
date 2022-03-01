package main

import "fmt"

func main() {
	m := map[string]int{
		"int1": 20,
		"int2": 21,
		"int3": 22,
		"int4": 23,
	}
	fmt.Println(m) // map[int1:20 int2:21 int3:22 int4:23]

	// append
	m["int5"] = 24
	m["int6"] = 25
	fmt.Println(m) // map[int1:20 int2:21 int3:22 int4:23 int5:24 int6:25]

	delete(m, "int1")
	delete(m, "int2")
	delete(m, "int3")
	fmt.Println(m) // map[int4:23 int5:24 int6:25]
}
