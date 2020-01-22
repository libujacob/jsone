package jsone

// Stack data structure.
type Stack struct {
	top    *entry
	length int
}

// Stack entry data structure.
type entry struct {
	stack *Stack
	next  *entry
	value interface{}
}

// NewStack creates a new Stack data structure object.
func NewStack() *Stack {
	return &Stack{top: nil, length: 0}
}

// Push a new entry to the stack.
func (s *Stack) Push(value interface{}) {
	s.top = &entry{stack: s, next: s.top, value: value}
	s.length++
}

// Pop out an entry from the stack.
func (s *Stack) Pop() interface{} {
	if s.length > 0 {
		value := s.top.value
		s.top = s.top.next
		s.length--
		return value
	}
	return nil
}

// Top gets the last inserted entry with out deleting from the stack.
func (s *Stack) Top() interface{} {
	if s.length > 0 {
		return s.top.value
	}
	return nil
}

// Size returns the number of the elements in the stack.
func (s *Stack) Size() int {
	return s.length
}
