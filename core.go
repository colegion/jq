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

// Element is a type that represents an element in the Document Object Model
// (DOM) that can have attributes, text, and children.
type Element struct {
}

// Q accepts a string containing a CSS selector and, optionally, a context
// and uses the selector to match a set of elements.
// Context is either a DOM Element or a jQuery object.
// This function is an equivalent of jQuery( selector [,context] )
// or $( selector [,context] ) in JavaScript.
func Q(selector string, context ...Context) *JQuery {
	return &JQuery{Global.Get("jQuery").New([]interface{}{selector, context})}
}
