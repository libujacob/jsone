package jsongo

import (
	"errors"
	"fmt"
	"reflect"
)

// O represents a Json Object.
type O map[string]interface{}

// Object creates a json object.
func Object() O {
	return O{}
}

// Put inserts an element into a json object.
func (jObj O) Put(key string, value interface{}) O {
	jObj[key] = value
	return jObj
}

// Get retrieves an element from a json object. Type of the return value is not predefined,
// caller has to check the return type.
func (jObj O) Get(key string) interface{} {
	return jObj[key]
}

// GetString retrieves a string data from a json object. Return error, if key not exist or data type not string.
func (jObj O) GetString(key string) (string, error) {
	switch jObj[key].(type) {
	case string:
		return jObj[key].(string), nil
	}
	return "", errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not string", key, reflect.TypeOf(jObj[key])))
}

// GetInt retrieves an int data from a json object. Return error, if key not exist or data type not int.
func (jObj O) GetInt(key string) (int, error) {
	switch jObj[key].(type) {
	case int:
		return jObj[key].(int), nil
	}

	return 0, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not int", key, reflect.TypeOf(jObj[key])))
}

// GetInt64 retrieves an int64 data from a json object. Return error, if key not exist or data type not int64.
func (jObj O) GetInt64(key string) (int64, error) {
	switch jObj[key].(type) {
	case int64:
		return jObj[key].(int64), nil
	}

	return 0, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not int64", key, reflect.TypeOf(jObj[key])))
}

// GetFloat64 retrieves a float64 data from a json object. Return error, if key not exist or data type not float64.
func (jObj O) GetFloat64(key string) (float64, error) {
	switch jObj[key].(type) {
	case float64:
		return jObj[key].(float64), nil
	}
	return 0.0, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not float64", key,
		reflect.TypeOf(jObj[key])))
}

// GetBoolean retrieves a boolean data from a json object. Return error, if key not exist or data type not boolean.
func (jObj O) GetBoolean(key string) (bool, error) {
	switch jObj[key].(type) {
	case bool:
		return jObj[key].(bool), nil
	}

	return false, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not boolean", key, reflect.TypeOf(jObj[key])))
}

// GetObject retrieves a json object data from a json object. Return error,
// if key not exist or data type not json object.
func (jObj O) GetObject(key string) (value O, err error) {
	switch jObj[key].(type) {
	case map[string]interface{}:
		object := Object()

		for k, v := range jObj[key].(map[string]interface{}) {
			object.Put(k, v)
		}

		return object, nil
	case O:
		return jObj[key].(O), nil
	}

	return nil, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not jsongo.object",
		key, reflect.TypeOf(jObj[key])))
}

// GetArray retrieves a json array data from a json object. Return error, if key not exist or data type not json array.
func (jObj O) GetArray(key string) (newArray *A, err error) {
	newArray = Array()

	switch jObj[key].(type) {
	case []interface{}:
		values := jObj[key].([]interface{})

		for _, value := range values {
			newArray.Put(value)
		}

		return newArray, nil
	case []string:
		values := jObj[key].([]string)

		for _, value := range values {
			newArray.Put(value)
		}

		return newArray, nil
	case *A:
		return jObj[key].(*A), nil
	}

	return nil, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not jsongo.A or []interface{}",
		key, reflect.TypeOf(jObj[key])))
}

// Remove an element from a json object.
func (jObj O) Remove(key string) O {
	delete(jObj, key)
	return jObj
}

// Has checks the object has an element in the name of the input string. Returns true if present, else false.
func (jObj O) Has(key string) bool {
	_, ok := jObj[key]
	if ok {
		return true
	}
	return false
}

// String on object generates a string representation of json object.
func (jObj O) String() string {
	return _string(jObj)
}

// Indent on object generates a string representation of json object with proper indent.
func (jObj O) Indent() string {
	return indent(jObj)
}
