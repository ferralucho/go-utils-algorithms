package main

import (
	"errors"
	"fmt"
	"sync"
)

type sliceStack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []int
}

func NewSliceStack() *sliceStack {
	return &sliceStack{sync.Mutex{}, make([]int, 0)}
}

func (s *sliceStack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *sliceStack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func main() {
	s := NewSliceStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
