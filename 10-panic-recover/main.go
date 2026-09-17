package main

import "fmt"

func riskyOperation() {
	fmt.Println("Starting risky operation")
	panic("something went catastrophically wrong")
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic", r)
		}

	}()

	riskyOperation()
	fmt.Println("this still won't run")
}

func main() {
	safeCall()
	fmt.Println("but main continues normally")
}
