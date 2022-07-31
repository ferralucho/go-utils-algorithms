package main

import "fmt"

type sliceStack struct {
	s []int32
}

func NewSliceStack() *sliceStack {
	return &sliceStack{make([]int32, 0)}
}

func (s *sliceStack) Push(v int32) {
	s.s = append(s.s, v)
}

func (s *sliceStack) GetItems() []int32 {
	return s.s
}

func (s *sliceStack) Pop() (int32, string) {

	l := len(s.s)
	if l == 0 {
		return 0, "error"
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, ""
}

func (s *sliceStack) Peek() int32 {

	l := len(s.s)
	if l == 0 {
		return 0
	}

	res := s.s[l-1]
	return res
}

func segment(x int32, space []int32) int32 {
	var cNum = int32(1)

	s := NewSliceStack()
	s.Push(0)

	for i := int32(0); i < int32(len(space)); i++ {
		if i < x {
			if space[i] < space[s.Peek()] {
				s.Pop()
				s.Push(i)
			}
		} else {
			var peekedItem = s.Peek()
			if peekedItem >= cNum {
				if space[i] < space[peekedItem] {
					s.Push(i)
				} else {
					s.Push(peekedItem)
				}
			} else {
				s.Push(i)
				var j = cNum

				for count := int32(0); count < x; count++ {
					if space[j] < space[s.Peek()] {
						s.Pop()
						s.Push(j)
					}
					j++
				}
			}
			cNum++
		}
	}

	fmt.Println("stack", *s)
	maxVal := int32(0)
	fmt.Println("space", space)

	for _, v := range s.GetItems() {
		support := space[v]
		if support > maxVal {
			maxVal = support
		}
	}

	return maxVal

}
