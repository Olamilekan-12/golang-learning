package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println(nums)

	nums[0] = 10
	nums[1] = 20

	fmt.Println(nums)

	scores := []int{10, 20, 30}
	fmt.Println(scores)
	fmt.Println(len(scores))

	scores = append(scores, 40)
	fmt.Println(scores)

	sub := scores[1:3]
	fmt.Println(sub)

	sub[0] = 999
	fmt.Println(sub)
	fmt.Println(scores)

	independent := make([]int, len(scores))
	copy(independent, scores)
	independent[0] = -1
	fmt.Println(independent)
	fmt.Println(scores)
	scores = append(scores, 50, 60, 70, 80)
	fmt.Println(scores)
	fmt.Println(independent)
	sub[0] = 500
	fmt.Println(sub)
	fmt.Println(scores)

}
