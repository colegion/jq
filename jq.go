// Package jq provides GopherJS bindings for jQuery, a fast, small,
// and feature-rich JavaScript library.
package jq

import (
	"github.com/goaltools/jq/internal"
)

// Global is a container for a native object that
// is used for interaction with the JavaScript API.
var Global internal.Object = internal.NewJS()
