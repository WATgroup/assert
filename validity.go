package assert

import (
	"testing"
)

func Noerror(t *testing.T, f func() error) {
	t.Helper()

	if e := f(); nil != e {
		t.Error(e.Error())
	}
}

func Valid(t *testing.T, x any) {

	if obj, ok := x.(interface{ Valid() bool }); !ok {
		t.Errorf("Provided value does not implement the `Valid() bool`  method")
	} else {

		if !obj.Valid() {
			t.Fail()
		}
	}
}
