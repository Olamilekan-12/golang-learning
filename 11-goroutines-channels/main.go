package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goroutine")
}

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("worker", i, "starting")
// 		}()
// 	}

//		wg.Wait()
//		fmt.Println("all workers done")
//	}
func main() {
	results := make(chan string, 3)

	for i := 1; i <= 3; i++ {
		go func(n int) {
			results <- fmt.Sprintf("worker %d finished", n)
		}(i)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
}
