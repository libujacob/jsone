package jsone

import (
	"fmt"
	"reflect"
)

// A A represents Json Array.
type A []interface{}

// Array creates a json array.
func Array() *A {
	return &A{}
}

// Put inserts an element into a json array.
func (jArray *A) Put(value interface{}) *A {
	*jArray = append(*jArray, value)
	return jArray
}

// String generates a string representation of a json array.
func (jArray *A) String() string {
	return _string(jArray)
}

// Indent generates a string of a json array with indent for formatting.
func (jArray *A) Indent() string {
	return indent(jArray)
}

// Size operation gets the number of elements in an array.
func (jArray *A) Size() int {
	return len(*jArray)
}

// OfString convert the json array into an array of strings.
func (jArray *A) OfString() (values []string, err error) {
	for _, value := range *jArray {
		if reflect.TypeOf(value).String() != "string" {
			return nil, fmt.Errorf("value is %s, not a string", reflect.TypeOf(value))
		}

		values = append(values, value.(string))
	}

	return values, nil
}
