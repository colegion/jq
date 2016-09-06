package jq

import (
	"github.com/colegion/jq/internal"
)

// Context is an
type Context interface {
	Object() interface{}
}

// JQuery is a type that represents a collection of Document Object Model
// (DOM) elements that have been created from an HTML string or selected
// from a document.
type JQuery struct {
	o internal.Object
}

// Object returns the native JS object. This method is
// necessary for implementation of the Context interface.
func (o *JQuery) Object() interface{} {
	return o.o
}

// Element is a type that represents an element in the Document Object Model
// (DOM) that can have attributes, text, and children.
type Element struct {
	o internal.Object
}

// Object returns the native JS object. This method is
// necessary for implementation of the Context interface.
func (o *Element) Object() interface{} {
	return o.o
}

// Q accepts a string containing a CSS selector and, optionally, a context
// and uses the selector to match a set of elements.
// Context is either a DOM Element or a jQuery object.
// If there are more than one context, only the first one is used and others
// are ignored.
// This function is an equivalent of jQuery( selector [,context] )
// or $( selector [,context] ) in JavaScript.
// Example 1:
//	Q("div.foo")
// If no element match the provided selector, the new JQuery object is "empty";
// that is, it contains no elements and has length of 0.
// Example 2:
//	Q("div.foo").Click(func(e *EventObject) {
//		// Restrict the search of the "span" to the
//		// clicked element.
//		Q("span", e).AddClass("bar")
//	})
func Q(selector string, context ...Context) *JQuery {
	if len(context) == 0 {
		return newJQuery([]interface{}{selector})
	}
	return newJQuery([]interface{}{selector, context[0].Object()})
}

// QElement gets a DOM element to wrap in a jQuery object.
// This function is an equivalent of jQuery( element )
// or $( element ) in JavaScript.
// Example:
//	Q("div.foo").Click(func(e *EventObject) {
//		// Hide the clicked element with a sliding animation.
//		QElement(e).SlideUp()
//	})
func QElement(element Element) *JQuery {
	return newJQuery([]interface{}{element})
}

// QElements gets a slice of DOM elements to wrap in a jQuery object.
// This function is an equivalent of jQuery( elements )
// or $( elements ) in JavaScript.
// Example:
//	Post("url.xml", func(d Data) {
//		child := QElements(d).Find("child")
//		_ = child
//	})
func QElements(elements []Element) *JQuery {
	return newJQuery([]interface{}{elements})
}

// QObject gets a plain object to wrap in a jQuery object.
// This function is an equivalent of jQuery( object )
// or $( object ) in JavaScipt where object is of type PlainObject.
// Operations supported on Plain objects wrapped in jQuery are:
// Data(), Prop(), On(), Off(), Trigger(), and TriggerHandler().
// The use of Data() on a plain object will result in a new property
// on the object called RandomNumber.
// Example:
//	// Define a plain object.
//	foo := struct {
//		Foo, Hello string
//	} {
//		"bar", "world",
//	}
//
//	// Pass it to the jQuery function.
//	obj := QObject(foo)
//
//	// Test accessing property values.
//	println(obj.Prop("Foo")) // bar
func QObject(object interface{}) *JQuery {
	return newJQuery([]interface{}{object})
}

// QJQuery clones and returns an existing jQuery object.
// This function is an equivalent of jQuery( selection )
// or $( selection ) in JavaScript.
func QJQuery(jq *JQuery) *JQuery {
	return newJQuery([]interface{}{jq})
}

// QEmpty allocates and returns a new empty jQuery object.
// This function is an equivalent of jQuery() or $() in JavaScript.
// jQuery 1.4+ is required.
func QEmpty() *JQuery {
	return newJQuery([]interface{}{})
}

// QFunc binds a function to be executed when the DOM has finished loading.
// This function is an equivalent of jQuery( callback ), $( callback )
// in JavaScript.
func QFunc(callback func()) *JQuery {
	return newJQuery([]interface{}{callback})
}

// newJQuery gets a slice of arguments, allocates a new JQuery
// object and returns it.
func newJQuery(args []interface{}) *JQuery {
	return &JQuery{Global.Get("jQuery").New(args)}
}
