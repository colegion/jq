package test

import (
	"github.com/colegion/jq/internal"
)

// JS implements an Object interface and is used for
// testing on non-JS platforms.
// Calls to its methods are recorded and can be
// retrieved using Test() method.
type JS struct {
	calls []Call
}

// Call is a type that represents information about a
// method call.
type Call struct {
	method string
	args   []interface{}
}

// NewJS allocated and returns a new JS object.
func NewJS() *JS {
	return &JS{}
}

// Get returns an Object containing information about
// all the previous calls and the current one.
func (o *JS) Get(key string) internal.Object {
	o.calls = append(o.calls, Call{"Get", []interface{}{key}})
	return o
}

// Call returns an Object containing information about
// all the previous calls and the current one.
func (o *JS) Call(name string, args []interface{}) internal.Object {
	o.calls = append(o.calls, Call{"Call", []interface{}{name, args}})
	return o
}

// Interface returns an Object containing information about
// all the previous calls and the current one.
func (o *JS) Interface() interface{} {
	o.calls = append(o.calls, Call{"Interface", []interface{}{}})
	return o
}

// New returns an Object containing information about
// all the previous calls and the current one.
func (o *JS) New(args []interface{}) internal.Object {
	o.calls = append(o.calls, Call{"New", []interface{}{args}})
	return o
}

// Test returns information about all the calls of the object
// that have ever been made.
func (o *JS) Test() interface{} {
	return o.calls
}
