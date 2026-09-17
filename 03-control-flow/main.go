package main

import "fmt"

func main() {
	age := 20

	if age >= 20 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}

	if score := 75; score >= 50 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Count : ", i)
	}

	n := 0

	for n < 3 {
		fmt.Println("n is ", n)
		n++
	}

	count := 0

	for {
		fmt.Println("looping")
		count++

		if count == 3 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("value: ", i)
	}

	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("unknown day")
	}

}
