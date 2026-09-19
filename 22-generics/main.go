package main

import (
	"cmp"
	"fmt"
)

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func main() {
	fmt.Println(Max(3, 7))
	fmt.Println(Max(2.5, 1.5))
	fmt.Println(Max("apple", "banana"))

	ints := &Stack[int]{}
	ints.Push(1)
	ints.Push(2)
	ints.Push(3)

	v, ok := ints.Pop()
	fmt.Println(v, ok)

	word := &Stack[string]{}
	word.Push("go")
	w, ok := word.Pop()
	fmt.Println(w, ok)

	_, ok = word.Pop()
	fmt.Println(ok)

}
