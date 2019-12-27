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
	switch this[key].(type) {
	case string:
		return this[key].(string), nil
	}
	return "", errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not string", key, reflect.TypeOf(this[key])))
}

func (this O) GetInt(key string) (int, error) {
	switch this[key].(type) {
	case int:
		return this[key].(int), nil
	}

	return 0, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not int", key, reflect.TypeOf(this[key])))
}

func (this O) GetInt64(key string) (int64, error) {
	switch this[key].(type) {
	case int64:
		return this[key].(int64), nil
	}

	return 0, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not int64", key, reflect.TypeOf(this[key])))
}

func (this O) GetFloat64(key string) (float64, error) {
	switch this[key].(type) {
	case float64:
		return this[key].(float64), nil
	}
	return 0.0, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not float64", key,
		reflect.TypeOf(this[key])))
}

func (this O) GetBoolean(key string) (bool, error) {
	switch this[key].(type) {
	case bool:
		return this[key].(bool), nil
	}

	return false, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not boolean", key, reflect.TypeOf(this[key])))
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

	return nil, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not jsongo.object",
		key, reflect.TypeOf(this[key])))
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

	return nil, errors.New(fmt.Sprintf("Casting error[%s]. Interface is %s, not jsongo.A or []interface{}",
		key, reflect.TypeOf(this[key])))
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
