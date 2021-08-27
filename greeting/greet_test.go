package greeting

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "André")

	got := buffer.String()
	want := "Hello, André"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}