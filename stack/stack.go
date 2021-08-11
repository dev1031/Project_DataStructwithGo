package stack

type s struct {
	array []int
	len   int
}

func (s *s) Push(elem int) {
	s.array = append(s.array, elem)
	s.len++
}

func (s *s) Pop() {
	if s.len != 0 {
		s.array = s.array[0 : s.len-1]
		s.len--
	}
}

func (s *s) Length() int {
	return s.len
}

// func (s *s) isEmpty() bool {
// 	return s.len == 0
// }

func Stack(arg int) []int {
	var stack s
	stack.Push(arg)
	// stack.Push(9)
	// stack.Push(10)
	// stack.Push(2)
	// stack.Push(5)
	// stack.Push(7)
	// stack.Push(8)
	// fmt.Println(stack.array)
	// fmt.Println(stack.len)
	// stack.Pop()
	// fmt.Println(stack.array, stack.len)
	// stack.Pop()
	// fmt.Println(stack.array, stack.len)
	// stack.Pop()
	// fmt.Println(stack.array, stack.len)
	// stack.Pop()
	// fmt.Println(stack.array, stack.len)
	// fmt.Println(stack.isEmpty())
	// fmt.Println(stack.Length())
	return stack.array
}
