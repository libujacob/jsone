JsonE
===================

**Single framework to parse and dynamically create/modify Json** objects.

[![godoc](https://godoc.org/github.com/libujacob/jsone?status.svg)](http://godoc.org/github.com/libujacob/jsone)
[![travis-ci](https://travis-ci.org/libujacob/jsone.svg)](https://travis-ci.org/libujacob/jsone) 
[![codecov](https://codecov.io/gh/libujacob/jsone/branch/master/graph/badge.svg)](https://codecov.io/gh/libujacob/jsone)
[![goreportcard](https://goreportcard.com/badge/github.com/libujacob/jsone)](http://gocover.io/github.com/libujacob/jsone)

Install
-------------
```
go get github.com/libujacob/jsone
```

Usage
-------------
### Create Json

To create this:  
```
{  
    "name":"Ricardo Longa",
    "idade":28,
    "owner":true,
    "skills":[  
        "Golang",
        "Android"
    ]
}
```  
Do this:  
```
import (
    j "github.com/libujacob/jsone"
)

json := j.Object().Put("name", "Ricardo Longa").
				   Put("idade", 28).
				   Put("owner", true).
				   Put("skills", j.Array().Put("Golang").
									       Put("Android"))

log.Println(json.Indent())
log.Println(json.String())
```
##### Convert object/array to indented String:
```
json.Indent()
```
##### Convert object/array to String:
```
json.String()
```
##### To remove a field of the object:
```
json.Remove("skills")
```
##### To get a field of the object:
```
json.Get("skills") // Return is interface{}.
```
##### To get string field of the object:
```
skill, err := json.GetString("skills") // Return is string, error
```
##### To get int/int64 field of the object:
```
count, err := json.GetInt("count") // Return is int, error
bytes, err := json.GetInt64("bytes") // Return is int64, error
```
##### To get float64 field of the object:
```
average, err := json.GetFloat64("average") // Return is float64, error
```
##### To get boolean field of the object:
```
isDownSupport, err := json.GetBoolean("isDownloadSupported") // Return is boolean, error
```
##### To check the object has a key:
```
if json.Has("operations") { // Return is boolean
    //do something
}
```
##### To range over a array:
```
results := Array().Put("Golang").Put("Android").Put("Java")

for i, result := range results.Array() {
}
```
##### To get Array size:
```
array := j.Array().Put("Android").
                   Put("Golang").
                   Put("Java")
                   
array.Size() // Result is 3.
```

### Parse Json
Json can be directly of Object or Array type. Both can be parsed using two 
different APIs which are mentioned below. After parsing you can use all the above 
said operations on the return value. 

##### Parse a Json Object string:
```
import (
    j "github.com/libujacob/jsone"
)

parsedObject := j.ParseJsonObject([]byte(`{"type": "oper", "nameList":["John", "Dan"], "id":205896}`))
/*{
    "type": "oper",
    "nameList": [
      "John",
      "Dan"
    ],
    "id": 205896
  }*/

parsedObject.Put("dept", "Operations")
/*{
    "type": "oper",
    "nameList": [
      "John",
      "Dan"
    ],
    "id": 205896,
    "dept": "Operations"
  }*/
```
##### Parse a Json Array string:
```
import (
    j "github.com/libujacob/jsone"
)

parsedArray := j.ParseJsonArray([]byte(`[{"name": "John", "id": 567314}, {"name": "Dan", "id": 589725}]`))
/*[
    {
      "name": "John",
      "id": 567314
    },
    {
      "name": "Dan",
      "id": 589725
    }
  ]*/

parsedArray.Put(j.Object().Put("name", "Tom").Put("id", 589289)
/*[
    {
      "name": "John",
      "id": 567314
    },
    {
      "name": "Dan",
      "id": 589725
    },
    {
      "name": "Tom",
      "id": 589289
    }
  ]*/
```

Copyright
-------------
Original work Copyright (c) 2015 Ricardo Longa.  
Modified work Copyright (c) 2019 Libu Jacob Varghese.  

Using bramp.net/antlr4/json and antlr for json parsing.

JsonE is licensed under the **Apache License Version 2.0**. See the LICENSE file for more information.
