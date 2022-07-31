package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Stack struct {
	dll   *list.List
	mutex sync.Mutex
}

func NewStack() *Stack {
	return &Stack{dll: list.New()}
}

func (s *Stack) Push(x interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.dll.PushBack(x)
}

func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.dll.Len() == 0 {
		return nil
	}
	tail := s.dll.Back()
	val := tail.Value
	s.dll.Remove(tail)
	return val
}

func main() {
	stack := NewStack()
	fmt.Println(stack.Pop())
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
