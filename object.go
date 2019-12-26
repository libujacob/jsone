package jsongo

import (
	"errors"
	"fmt"
	"reflect"
)

type O map[string]interface{}

func Object() O {
	return O{}
}

func (this O) Put(key string, value interface{}) O {
	this[key] = value
	return this
}

func (this O) Get(key string) interface{} {
	return this[key]
}

func (this O) GetString(key string) (string, error) {
	if reflect.TypeOf(this[key]).Kind() == reflect.String {
		return this[key].(string), nil
	}
	return "", &JsonError{op: "GetString", element: key, msg: "type miss-match."}
}

func (this O) GetInt(key string) (int, error) {
	if reflect.TypeOf(this[key]).Kind() == reflect.Int {
		return this[key].(int), nil
	}
	return 0, &JsonError{op: "GetInt", element: key, msg: "type miss-match."}
}

func (this O) GetFloat64(key string) (float64, error) {
	if reflect.TypeOf(this[key]).Kind() == reflect.Float64 {
		return this[key].(float64), nil
	}
	return 0.0, &JsonError{op: "GetFloat64", element: key, msg: "type miss-match."}
}

func (this O) GetBoolean(key string) (bool, error) {
	if reflect.TypeOf(this[key]).Kind() == reflect.Bool {
		return this[key].(bool), nil
	}
	return false, &JsonError{op: "GetBoolean", element: key, msg: "type miss-match."}
}

func (this O) GetObject(key string) (value O, err error) {
	switch this[key].(type) {
	case map[string]interface{}:
		object := Object()

		for k, v := range this[key].(map[string]interface{}) {
			object.Put(k, v)
		}

		return object, nil
	case O:
		return this[key].(O), nil
	}

	return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.object", reflect.TypeOf(this[key])))
}

func (this O) GetArray(key string) (newArray *A, err error) {
	newArray = Array()

	switch this[key].(type) {
	case []interface{}:
		values := this[key].([]interface{})

		for _, value := range values {
			newArray.Put(value)
		}

		return newArray, nil
	case []string:
		values := this[key].([]string)

		for _, value := range values {
			newArray.Put(value)
		}

		return newArray, nil
	case *A:
		return this[key].(*A), nil
	}

	return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.A or []interface{}", reflect.TypeOf(this[key])))
}

func (this O) Remove(key string) O {
	delete(this, key)
	return this
}

func (this O) Indent() string {
	return indent(this)
}

func (this O) String() string {
	return _string(this)
}
