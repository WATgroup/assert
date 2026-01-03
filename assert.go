// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2 OR Proprietary
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

// Package assert provides various simple assertion primitives, for use with "go test"
package assert

import (
	"testing"
)

func Equal[V comparable](t *testing.T, got, expected V) {
	t.Helper()

	if expected != got {
		t.Errorf(`assert·Equal(got: %v, expected: %v)`, got, expected)
	}
}

func True(t *testing.T, value bool) {
	t.Helper()

	if !value {
		t.Errorf(`assert·True(%v)`, value)
	}
}

func False(t *testing.T, value bool) {
	t.Helper()

	if value {
		t.Errorf(`assert·False(%v)`, value)
	}
}
