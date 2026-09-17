package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	fmt.Println(ages)
	fmt.Println(ages["Alice"])

	value, ok := ages["Alice"]
	fmt.Println(value, ok)

	value2, ok2 := ages["Charlie"]
	fmt.Println(value2, ok2)

	ages["Charlie"] = 40
	fmt.Println(ages)

	ages["Alice"] = 31
	fmt.Println(ages)

	delete(ages, "Bob")
	fmt.Println(ages)

	for key, value := range ages {
		fmt.Println(key, "->", value)
	}

}
