package jsone

import (
	"bytes"
	"strings"
	"testing"
)

func Test_create_empty_object(t *testing.T) {
	expect := bytes2json([]byte(`{}`))
	result := Object()

	check(t, struct2json(expect), struct2json(result))
}

func Test_create_populated_object(t *testing.T) {
	expect := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28,"owner":true,"skills":["Golang","Android"]}`))
	result := Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("owner", true).Put("skills", Array().Put("Golang").Put("Android"))

	check(t, struct2json(expect), struct2json(result))
}

func Test_create_populated_objects_and_remove_attr(t *testing.T) {
	expect := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28,"skills":["Golang","Android"]}`))
	result := Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("skills", Array().Put("Golang").Put("Android"))

	check(t, struct2json(expect), struct2json(result))

	expectAfterRemove := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28}`))

	result.Remove("skills")

	check(t, struct2json(expectAfterRemove), struct2json(result))
}

func Test_object_get_func(t *testing.T) {
	expect := "Ricardo Longa"
	result := Object().Put("name", "Ricardo Longa")

	if !strings.EqualFold(expect, result.Get("name").(string)) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, result.Get("name"))
	}
}

func Test_object_indent(t *testing.T) {
	expect := []byte(`{
   "skills": [
      "Golang",
      "Android",
      "Java"
   ]
}`)
	result := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	if !bytes.Equal(expect, bytes.NewBufferString(result.Indent()).Bytes()) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, struct2json(result.Indent()))
	}
}

func Test_object_string(t *testing.T) {
	expect := []byte(`{"skills":["Golang","Android","Java"]}`)
	result := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	if !bytes.Equal(expect, bytes.NewBufferString(result.String()).Bytes()) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, struct2json(result.String()))
	}
}

func Test_get_object_with_casting_error(t *testing.T) {
	obj := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	if _, err := obj.GetObject("skills"); err == nil {
		t.Errorf("Casting error not found.")
	}
}

func Test_get_object_without_casting_error(t *testing.T) {
	obj := Object().Put("owner", Object().Put("nome", "Ricardo Longa"))

	_, err := obj.GetObject("owner")
	if err != nil {
		t.Errorf("1Casting error not expected.")
	}

	obj = Object().Put("owner", map[string]interface{}{
		"nome": "Ricardo Longa",
	})

	_, err = obj.GetObject("owner")
	if err != nil {
		t.Errorf("2Casting error not expected.")
	}
}

func Test_get_array_without_casting_error(t *testing.T) {
	obj := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	values, err := obj.GetArray("skills")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if len(*values) != 3 {
		t.Error("Expected 3 values.")
	}

	obj = Object().Put("skills", []interface{}{"Golang", "Android", "Java"})

	values, err = obj.GetArray("skills")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if len(*values) != 3 {
		t.Error("Expected 3 values.")
	}

	obj = Object().Put("skills", []string{"Golang", "Android", "Java"})

	values, err = obj.GetArray("skills")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if len(*values) != 3 {
		t.Error("Expected 3 values.")
	}
}

func Test_get_array_with_casting_error(t *testing.T) {
	obj := Object().Put("owner", Object().Put("nome", "Ricardo Longa"))

	if _, err := obj.GetArray("owner"); err == nil {
		t.Errorf("Casting error not found.")
	}
}

func Test_has_key_exists(t *testing.T) {
	obj := Object().Put("owner", Object().Put("name", "Jacob Varghese"))

	if !obj.Has("owner") {
		t.Errorf("Has key error, key not exists.")
	}
}

func Test_has_key_not_exists(t *testing.T) {
	obj := Object().Put("owner", Object().Put("name", "Jacob Varghese"))

	if obj.Has("operations") {
		t.Errorf("Has key error, key exists.")
	}
}

func Test_get_string(t *testing.T) {
	name := "Jacob Varghese"
	obj := Object().Put("owner", Object().Put("name", name).Put("experience", 23))

	owner, err := obj.GetObject("owner")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	ownerName, err := owner.GetString("name")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if name != ownerName {
		t.Errorf("Value extraction missmatch:")
	}

	_, err = owner.GetString("experience")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}

	_, err = owner.GetString("dept")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}
}

func Test_get_int(t *testing.T) {
	age := 35
	obj := Object().Put("owner", Object().Put("age", age).Put("experience", 12.4))

	owner, err := obj.GetObject("owner")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	ownerAge, err := owner.GetInt("age")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if age != ownerAge {
		t.Errorf("Value extraction missmatch:")
	}

	_, err = owner.GetInt("experience")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}

	_, err = owner.GetInt("dept")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}

	_, err = owner.GetInt64("age")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}
}

func Test_get_float64(t *testing.T) {
	experience := 12.8
	obj := Object().Put("owner", Object().Put("name", "Jacob Varghese").Put("experience", experience))

	owner, err := obj.GetObject("owner")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	ownerExperience, err := owner.GetFloat64("experience")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if experience != ownerExperience {
		t.Errorf("Value extraction missmatch:")
	}

	_, err = owner.GetFloat64("name")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}

	_, err = owner.GetFloat64("dept")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}
}

func Test_get_boolean(t *testing.T) {
	inRolls := true
	obj := Object().Put("owner", Object().Put("name", "Jacob Varghese").Put("inRolls", inRolls))

	owner, err := obj.GetObject("owner")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	ownerInRolls, err := owner.GetBoolean("inRolls")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if inRolls != ownerInRolls {
		t.Errorf("Value extraction missmatch:")
	}

	_, err = owner.GetBoolean("name")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}

	_, err = owner.GetBoolean("dept")
	if err == nil {
		t.Errorf("Error expected but not happened: %s.", err)
	}
}

func Test_Keys(t *testing.T) {
	obj := Object().Put("owner", Object().Put("name", "Jacob Varghese").Put("inRolls", true))

	// Case1: One key
	check(t, []string{"owner"}, obj.Keys())

	// Case2: 2 Keys
	owner, _ := obj.GetObject("owner")
	fieldCount := 0
	for _, k := range owner.Keys() {
		if k != "name" && k != "inRolls" {
			t.Errorf("Invalid entry in the return list: %s.", k)
		} else {
			fieldCount++
		}
	}
	check(t, fieldCount, 2)

	// Case3: Nil input and empty key
	nilSample, _ := obj.GetObject("sample")
	check(t, []string{}, nilSample.Keys())

}
