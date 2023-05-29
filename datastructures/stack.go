package datastructures

type stack struct {
	items []string
}

func NewStack() *stack {
	items := []string{}
	return &stack{items: items}
}

func (s *stack) Push(v string) {
	s.items = append(s.items, v)
}

func (s *stack) Pop() string {
	i := len(s.items) - 1

	top := s.items[i]
	s.items = s.items[:i]

	return top
}

func (s *stack) Size() int {
	return len(s.items)
}
