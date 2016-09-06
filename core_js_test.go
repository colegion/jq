// +build js

package jq

import (
	"testing"
)

func TestQueries(t *testing.T) {
}

var testParentHTML = `<html>
	<head>
		<title>Foo Bar</title>
	</head>

	<body>
	</body>
</html>`

var testChildHTML = `<div class="foo">
	<span>Hello, world!</span>
</div>`
