package jsone

import (
	"testing"
)

func Test_stack(t *testing.T) {
	stack := NewStack()
	check(t, nil, stack.Top())
	check(t, 0, stack.Size())

	intVal := 2
	stack.Push(intVal)
	check(t, intVal, stack.Top())
	check(t, 1, stack.Size())

	stringVal := "hello"
	stack.Push(stringVal)
	check(t, stringVal, stack.Top())
	check(t, 2, stack.Size())

	popValue1 := stack.Pop()
	check(t, stringVal, popValue1)
	check(t, 1, stack.Size())

	popValue2 := stack.Pop()
	check(t, intVal, popValue2)
	check(t, 0, stack.Size())
	check(t, nil, stack.Top())

	popValue3 := stack.Pop()
	check(t, nil, popValue3)
	check(t, 0, stack.Size())
	check(t, nil, stack.Top())
}
