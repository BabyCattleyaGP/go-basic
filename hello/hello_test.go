package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to person", func(t *testing.T) {
		got := Hello("Baby", "")
		want := "Hello, Baby"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' if string empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Baby", "french")
		want := "Bonjour, Baby"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Baby", "spanish")
		want := "Hola, Baby"
		assertCorrectMessage(t, got, want)
	})
}
