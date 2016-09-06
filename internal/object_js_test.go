// +build js

package internal

import (
	"testing"

	"github.com/gopherjs/gopherjs/js"
)

func TestJS(t *testing.T) {
	g := &JS{js.Global}
	testObj := g.Call("eval", []interface{}{`({
		add: function(a, b) {
			return a + b;
		},
	})`})

	add := testObj.Get("add").Interface().(func(...interface{}) *js.Object)
	var i int64 = 40
	j := 2
	exp := int(i) + j
	if res := add(i, j).Int(); res != exp {
		t.Errorf(`Incorrect result. Expected: %d, got: %d.`, exp, res)
	}
}
