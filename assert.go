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
