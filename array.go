package jsongo

import (
	"errors"
	"fmt"
	"reflect"
)

// Json Array.
type A []interface{}

// Create a json array.
func Array() *A {
	return &A{}
}

// Insert an element into a json array.
func (jArray *A) Put(value interface{}) *A {
	*jArray = append(*jArray, value)
	return jArray
}

// Generate a string representation of a json array.
func (jArray *A) String() string {
	return _string(jArray)
}

// Generate a string of a json array with indent for formatting.
func (jArray *A) Indent() string {
	return indent(jArray)
}

// To get the number of elements in an array.
func (jArray *A) Size() int {
	return len(*jArray)
}

// Convert the json array into an array of strings.
func (jArray *A) OfString() (values []string, err error) {
	for _, value := range *jArray {
		if reflect.TypeOf(value).String() != "string" {
			return nil, errors.New(fmt.Sprintf("Value is %s, not a string.", reflect.TypeOf(value)))
		}

		values = append(values, value.(string))
	}

	return values, nil
}
