package test

import (
	"reflect"
	"testing"

	"github.com/colegion/jq/internal"
)

const incorrectCalls = "Incorrect list of calls. Expected %#v, got %#v."

func TestJS_NoCalls(t *testing.T) {
	o := Global
	var exp []Call
	if cs := o.Test(); !reflect.DeepEqual(cs.([]Call), exp) {
		t.Errorf(incorrectCalls, exp, cs)
	}
}

func TestJS(t *testing.T) {
	arg1 := "xxx"
	o := Global.Get(arg1)

	arg2 := "eval"
	arg3 := "x1"
	arg4 := "x2"
	o = o.Call(arg2, []interface{}{arg3, arg4})

	arg5 := "a1"
	arg6 := "a2"
	o = o.New([]interface{}{arg5, arg6})

	o = o.Interface().(internal.Object)

	exp := []Call{
		Call{"Get", []interface{}{arg1}},
		Call{"Call", []interface{}{arg2, []interface{}{arg3, arg4}}},
		Call{"New", []interface{}{[]interface{}{arg5, arg6}}},
		Call{"Interface", []interface{}{}},
	}
	if cs := o.Test(); !reflect.DeepEqual(cs, exp) {
		t.Errorf(incorrectCalls, exp, cs)
	}
}

var Global = NewJS()
