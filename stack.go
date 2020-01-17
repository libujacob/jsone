package jsongo

// Stack data structure
type Stack struct {
	top *entry
	len int
}

// Stack entry
type entry struct {
	stack *Stack
	next  *entry
	value interface{}
}

// NewStack create a new Stack data structure object
func NewStack() *Stack {
	return &Stack{top: nil, len: 0}
}

// Push new entry
func (s *Stack) Push(value interface{}) {
	s.top = &entry{stack: s, next: s.top, value: value}
	s.len++
}

// Pop out an entry
func (s *Stack) Pop() interface{} {
	if s.len > 0 {
		value := s.top.value
		s.top = s.top.next
		s.len--
		return value
	}
	return nil
}

// Top to get the last inserted an entry
func (s *Stack) Top() interface{} {
	if s.len > 0 {
		return s.top.value
	}
	return nil
}
