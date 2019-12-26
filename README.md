Jsongo
===================

**Fluent API** to make it easier **to parse and create Json** objects.

This **single framework** allows you to **parse and dynamically create/modify
Json** in GoLang.

[![travis-ci](https://travis-ci.org/ricardolonga/jsongo.svg)](https://travis-ci.org/ricardolonga/jsongo) 
[![codecov](https://codecov.io/gh/ricardolonga/jsongo/branch/master/graph/badge.svg)](https://codecov.io/gh/ricardolonga/jsongo)
[![goreportcard](https://goreportcard.com/badge/github.com/ricardolonga/jsongo)](http://gocover.io/github.com/ricardolonga/jsongo)

Install
-------------
```
go get github.com/libujacob/jsongo
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
    j "github.com/libujacob/jsongo"
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
    j "github.com/libujacob/jsongo"
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
    j "github.com/libujacob/jsongo"
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

Jsongo is licensed under the **Apache License Version 2.0**. See the LICENSE file for more information.
