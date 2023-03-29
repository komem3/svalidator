package svalidator_test

import (
	"errors"
	"testing"
)

func pointer[T any](t T) *T {
	return &t
}

func assertError(t *testing.T, want, got error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("want: %v.\nbut got: %v", want, got)
	}
}

func assertIsError(t *testing.T, want bool, got error) {
	t.Helper()
	if gotErr := got != nil; want != gotErr {
		t.Errorf("want: %v.\nbut got: %v", want, got)
	}
}
