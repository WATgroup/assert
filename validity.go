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
