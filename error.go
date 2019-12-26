package jsongo

import "fmt"

type JsonError struct {
	op      string // Operation
	msg     string // description of error
	element string // error occurred after reading Offset bytes
}

func (e *JsonError) Error() string {
	return fmt.Sprintf("Json %s error: [%s] %s", e.op, e.element, e.msg)
}
