// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2 OR Proprietary
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

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
