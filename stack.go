package jsongo

// Stack
type stack struct {
	top *entry
	len int
}

// Stack entry
type entry struct {
	stack *stack
	next  *entry
	value interface{}
}

// NewStack create a new stack data structure object
func NewStack() *stack {
	return &stack{top: nil, len: 0}
}

// Push new entry
func (s *stack) Push(value interface{}) {
	s.top = &entry{stack: s, next: s.top, value: value}
	s.len++
}

// POP out an entry
func (s *stack) Pop() interface{} {
	if s.len > 0 {
		value := s.top.value
		s.top = s.top.next
		s.len--
		return value
	}
	return nil
}

// POP out an entry
func (s *stack) Top() interface{} {
	if s.len > 0 {
		return s.top.value
	}
	return nil
}
