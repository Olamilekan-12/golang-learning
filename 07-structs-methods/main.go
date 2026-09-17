package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

func (p *Person) HaveBirthday() {
	p.Age = p.Age + 1
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println(p)
	fmt.Println(p.Name, p.Age)

	var empty Person
	fmt.Println(empty)

	fmt.Println(p.Greet())
	p.HaveBirthday()
	fmt.Println(p.Age)

	x := 10
	fmt.Println(x)
	fmt.Println(&x)

	ptr := &x
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 20
	fmt.Println(x)
}
