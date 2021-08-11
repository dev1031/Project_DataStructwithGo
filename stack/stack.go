package stack

import "fmt"

type s struct {
	array []int
	len   int
}

func (s *s) Add(elem int) {
	s.array = append(s.array, elem)
	s.len++
}

func Stack() {
	var stack s
	stack.Add(1)
	stack.Add(9)
	stack.Add(10)
	stack.Add(2)
	stack.Add(5)
	stack.Add(7)
	stack.Add(8)
	fmt.Println(stack.array)
	fmt.Println(stack.len)
}
