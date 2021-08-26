package main

import "testing"

/*
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
 */

/*
func TestHello(t *testing.T) {
	got := Hello("André")
	want := "Hello, André"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
 */

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("André", "")
		want := "Hello, André"
		assertCorrectMessage(t, got, want)
		})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Eduardo", "Spanish")
		want := "Hola, Eduardo"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Eduardo", "French")
		want := "Bonjour, Eduardo"
		assertCorrectMessage(t, got, want)
	})
}