package jsongo

import (
	"log"
	"strconv"

	"bramp.net/antlr4/json"                    // The parser
	"github.com/antlr/antlr4/runtime/Go/antlr" // The antlr library
)

// jsonListener is an event-driven callback for the parser.
type jsonListener struct {
	*json.BaseJSONListener // https://godoc.org/bramp.net/antlr4/json#BaseJSONListener
	jsonObject             O
	jsonArray              *A
	jsonEntryStack         *stack
}

func (l *jsonListener) EnterJson(ctx *json.JsonContext) {
	l.jsonEntryStack = NewStack()
}

func (l *jsonListener) ExitJson(ctx *json.JsonContext) {
}

func (l *jsonListener) EnterObj(ctx *json.ObjContext) {
	if l.jsonEntryStack.len == 0 {
		l.jsonObject = Object()
		l.jsonEntryStack.Push(l.jsonObject)
	} else {
		l.jsonEntryStack.Push(Object())
	}
}

func (l *jsonListener) ExitObj(ctx *json.ObjContext) {
	// Object can be placed in a pair or on an array itself, so here we handle
	// both the cases
	value := l.jsonEntryStack.Pop()

	// Popping the out the key from the stack to fill the object
	key := getKeyFromStack(l)
	insertIntoStackTopValue(l, key, value)
}

func (l *jsonListener) EnterPair(ctx *json.PairContext) {
	key := ctx.STRING().GetText()
	l.jsonEntryStack.Push(key[1 : len(key)-1])
}

func (l *jsonListener) ExitPair(ctx *json.PairContext) {
}

func (l *jsonListener) EnterArray(ctx *json.ArrayContext) {
	if l.jsonEntryStack.len == 0 {
		l.jsonArray = Array()
		l.jsonEntryStack.Push(l.jsonArray)
	} else {
		l.jsonEntryStack.Push(Array())
	}
}

func (l *jsonListener) ExitArray(ctx *json.ArrayContext) {
	// Arrays can be in a pair on on another array itself. So here we handle
	// both cases.
	value := l.jsonEntryStack.Pop()

	// Popping the out the key from the stack to fill the object
	key := getKeyFromStack(l)

	insertIntoStackTopValue(l, key, value)
}

func (l *jsonListener) EnterValue(ctx *json.ValueContext) {
}

func (l *jsonListener) ExitValue(ctx *json.ValueContext) {
	// Popping the out the key from the stack to fill the object
	key := getKeyFromStack(l)

	if ctx.Array() != nil || ctx.Obj() != nil {
		return
	}

	// Handle only terminal values here
	if nil != ctx.STRING() {
		value := ctx.STRING().GetText()
		value = value[1 : len(value)-1]
		insertIntoStackTopValue(l, key, value)
	} else if nil != ctx.NUMBER() {
		value := ctx.NUMBER().GetText()
		if !isNumericValue(value) {
			log.Panicf("Invalid numeric value [%s]!", value)
		}
		if isFloatValue(value) {
			numValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Panicf("Invalid numeric value [%s], parsing failed!", value)
			}
			insertIntoStackTopValue(l, key, numValue)
		} else {
			if len(value) > 1 && value[0:1] == "0" {
				log.Panicf("Invalid numeric value [%s], only decimal is supported!", value)
			}
			numValue, err := strconv.ParseInt(value, 0, 64)
			if err != nil {
				log.Panicf("Invalid numeric value [%s], parsing failed!", value)
			}
			insertIntoStackTopValue(l, key, numValue)
		}
	} else {
		switch ctx.GetText() {
		case "true":
			{
				insertIntoStackTopValue(l, key, true)
				break
			}
		case "false":
			{
				insertIntoStackTopValue(l, key, false)
				break
			}
		case "null":
			{
				insertIntoStackTopValue(l, key, nil)
				break
			}
		default:
			log.Panic("Invalid json string!")
		}
		return
	}
}

// Get next key from the stack
func getKeyFromStack(l *jsonListener) string {
	key := ""
	// Popping the out the key from the stack to fill the object
	switch l.jsonEntryStack.Top().(type) {
	case string:
		last := l.jsonEntryStack.Pop()
		if last == nil {
			log.Panic("No data in stack!")
		}
		key = last.(string)
	}
	return key
}

// Insert the value in to the next object into the stack
func insertIntoStackTopValue(l *jsonListener, key string, value interface{}) {
	switch l.jsonEntryStack.Top().(type) {
	case O:
		if key != "" {
			l.jsonEntryStack.Top().(O).Put(key, value)
		} else {
			log.Panic("No key string in stack to insert the array value!")
		}
		break

	case *A:
		l.jsonEntryStack.Top().(*A).Put(value)
		break

	default:
		// log.Print("Not a proper data type expected here")
		break
	}
}

// ParseJsonObject parse the json object to generate the json map structure
func ParseJsonObject(inData []byte) O {
	// Setup the input
	is := antlr.NewInputStream(string(inData))
	// Create the Lexer
	lexer := json.NewJSONLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := json.NewJSONParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// Finally walk the tree
	tree := p.Json()
	listener := jsonListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	return listener.jsonObject
}

// ParseJsonArray parse an array to generate the json map structure
func ParseJsonArray(inData []byte) *A {
	// Setup the input
	is := antlr.NewInputStream(string(inData))
	// Create the Lexer
	lexer := json.NewJSONLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := json.NewJSONParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// Finally walk the tree
	tree := p.Json()
	listener := jsonListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	return listener.jsonArray
}

// Check the number value is correct
func isNumericValue(s string) bool {
	containsDot := 0
	for _, ch := range s {
		if ch == '.' {
			containsDot++
		} else if ch < '0' || ch > '9' {
			return false
		}
	}

	return containsDot == 0 || containsDot == 1
}

// Check whether the number is float or not
func isFloatValue(s string) bool {
	for _, ch := range s {
		if ch == '.' {
			return true
		}
	}
	return false
}
