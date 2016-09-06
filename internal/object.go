package internal

import (
	"github.com/gopherjs/gopherjs/js"
)

// Object is an interface that must be implemented
// to be able to interact with the JavaScript API.
type Object interface {
	Get(string) Object
	Call(string, []interface{}) Object
	Interface() interface{}
	New([]interface{}) Object

	// Test should be a method for internal use
	// by tests only.
	Test() interface{}
}

// JS is a container type that represents a native JavaScript
// object. Calls to its methods are treated specially by GopherJS
// and translated directly to the JS.
// It implements Object interface.
type JS struct {
	object *js.Object
}

// NewJS allocates and returns a new instance of JS.
func NewJS() *JS {
	return &JS{js.Global}
}

// Get returns the object's property with the given key.
func (o *JS) Get(key string) Object {
	return &JS{o.object.Get(key)}
}

// Call gets a method name and, optionally, a number of arguments
// and calls the object's method using the parameters.
func (o *JS) Call(name string, args []interface{}) Object {
	return &JS{o.object.Call(name, args...)}
}

// New gets a number of input arguments and allocates a new instance
// of the this type object. If the object is not a constructor,
// this function will fail.
func (o *JS) New(args []interface{}) Object {
	return &JS{o.object.New(args...)}
}

// Interface returns an interface{} representation of the object.
func (o *JS) Interface() interface{} {
	return o.object.Interface()
}

// Test is a dummy method that is necessary for implementation of
// the Object interface. It may be useful for testing purposes.
// For details, see ./test package.
func (o *JS) Test() interface{} {
	return nil
}
