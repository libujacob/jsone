package jsone

// Stack data structure
type Stack struct {
	top    *entry
	length int
}

// Stack entry
type entry struct {
	stack *Stack
	next  *entry
	value interface{}
}

// NewStack create a new Stack data structure object
func NewStack() *Stack {
	return &Stack{top: nil, length: 0}
}

// Push new entry
func (s *Stack) Push(value interface{}) {
	s.top = &entry{stack: s, next: s.top, value: value}
	s.length++
}

// Pop out an entry
func (s *Stack) Pop() interface{} {
	if s.length > 0 {
		value := s.top.value
		s.top = s.top.next
		s.length--
		return value
	}
	return nil
}

// Top to get the last inserted an entry
func (s *Stack) Top() interface{} {
	if s.length > 0 {
		return s.top.value
	}
	return nil
}

// Size returns the length of the stack
func (s *Stack) Size() int {
	return s.length
}
