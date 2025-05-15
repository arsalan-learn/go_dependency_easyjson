package main

import (
	"fmt"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// This file demonstrates importing and using easyjson's utilities
// without requiring the generated code

// DemoEasyJSONUsage shows how to use easyjson packages directly
func DemoEasyJSONUsage() {
	// Create a jwriter (used by easyjson internally)
	w := jwriter.Writer{}
	w.RawString(`{"name":"John"}`)

	// Example of using jlexer (used by easyjson internally)
	data := []byte(`{"name":"John","age":30}`)
	lexer := jlexer.Lexer{Data: data}
	lexer.FetchToken()

	// Just for demonstration
	fmt.Println("Using easyjson utilities directly")
}
