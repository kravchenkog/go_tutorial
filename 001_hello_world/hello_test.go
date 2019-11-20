package hello_test

import "testing"

func TEstMe(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Error("Hello() = %q", got, want)
	}
}
