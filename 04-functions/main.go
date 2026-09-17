package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func divideNamed(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func sum(numbers ...int) (total int) {
	for _, n := range numbers {
		total += n
	}
	return
}

func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func demoDefer() {
	fmt.Println("start")
	defer fmt.Println("deferred call")
	defer fmt.Println("second deferred call")
	fmt.Println("end")
}

func main() {
	result := add(3, 4)
	fmt.Println(result)

	q, r := divide(17, 5)
	fmt.Println(q, r)

	q2, r2 := divideNamed(20, 6)
	fmt.Println(q2, r2)

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
	counter := makeCounter()
	counter2 := makeCounter()
	fmt.Println(counter2())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	demoDefer()
}
