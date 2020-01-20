package jsone

import (
	"testing"
)

func Test_Parsing_Object_Success_Cases(t *testing.T) {
	inputJson := `{"example": "json"}`
	expected := `{"example":"json"}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example": null}`
	expected = `{"example":null}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example": true}`
	expected = `{"example":true}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example": false}`
	expected = `{"example":false}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example":"json", "with":["an", "array"], "a":2}`
	expected = `{"a":2,"example":"json","with":["an","array"]}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example":"json", "with":["an", "array"], "an": [{"a": 1, "b": "text"}, {"c": true, "d": false}]}`
	expected = `{"an":[{"a":1,"b":"text"},{"c":true,"d":false}],"example":"json","with":["an","array"]}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example":"json", "with":["an", "array"], "an": [[1, 2, 3],[4, 5, 6, 7]]}`
	expected = `{"an":[[1,2,3],[4,5,6,7]],"example":"json","with":["an","array"]}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example":"json", "an": [[1, 2], 4, [5, [6, [7], [8], [[9]]]]]}`
	expected = `{"an":[[1,2],4,[5,[6,[7],[8],[[9]]]]],"example":"json"}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example": {"a": {"b": {"c": {"d":{"e": null}}}}}}`
	expected = `{"example":{"a":{"b":{"c":{"d":{"e":null}}}}}}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example": {"a": {"b": {"c": {"d": {"e": null}, "f": false}, "g": 1}, "h": "hello"}, "i": 5.322}, 
"j": 0}`
	expected = `{"example":{"a":{"b":{"c":{"d":{"e":null},"f":false},"g":1},"h":"hello"},"i":5.322},"j":0}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

	inputJson = `{"example": 0.1}`
	expected = `{"example":0.1}`
	check(t, expected, ParseJsonObject([]byte(inputJson)).String())

}

func Test_Parsing_Object_Failure_Case_0(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Value \"nil\" should have panicked!")
			}
		}()
		// This function should cause a panic
		ParseJsonObject([]byte(`{"example": nil}`))
	}()
}

func Test_Parsing_Array_Success_Cases(t *testing.T) {
	inputJson := `[{"example": "json"}]`
	expected := `[{"example":"json"}]`
	check(t, expected, ParseJsonArray([]byte(inputJson)).String())

	inputJson = `[{"example": "json"}, 1]`
	expected = `[{"example":"json"},1]`
	check(t, expected, ParseJsonArray([]byte(inputJson)).String())

	inputJson = `[null, {"example": "json"}]`
	expected = `[null,{"example":"json"}]`
	check(t, expected, ParseJsonArray([]byte(inputJson)).String())

	inputJson = `[{"example": {"a": {"b": {"c": {"d": {"e": null}, "f": false}, "g": 1}, "h": "hello"}, "i": 5.322}, 
"j": 0}, {"example": {"a": {"b": {"c": {"d": {"e": null}, "f": false}, "g": 1}, "h": "hello"}, "i": 5.322}, 
"j": 0}, {"example": {"a": {"b": {"c": {"d": {"e": null}, "f": false}, "g": 1}, "h": "hello"}, "i": 5.322}, 
"j": 0}]`
	expected = `[{"example":{"a":{"b":{"c":{"d":{"e":null},"f":false},"g":1},"h":"hello"},"i":5.322},"j":0},{"example":{"a":{"b":{"c":{"d":{"e":null},"f":false},"g":1},"h":"hello"},"i":5.322},"j":0},{"example":{"a":{"b":{"c":{"d":{"e":null},"f":false},"g":1},"h":"hello"},"i":5.322},"j":0}]`
	check(t, expected, ParseJsonArray([]byte(inputJson)).String())

	inputJson = `[{"a":["b", {"c": [5678090, "test", {"d":[1,2,3,4,5]}]}]}, {"example": "json"}]`
	expected = `[{"a":["b",{"c":[5678090,"test",{"d":[1,2,3,4,5]}]}]},{"example":"json"}]`
	check(t, expected, ParseJsonArray([]byte(inputJson)).String())
}
