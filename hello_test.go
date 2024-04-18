package main

import "testing"

// Tests should have CLEAR specifications

func TestHelloWithName(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := HelloWithName("Carlos", "English")
		want := "Hello, Carlos"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := HelloWithName("", "English")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Fernanda is my beloved's name, so you can't use it", func(t *testing.T) {
		got := HelloWithName("Fernanda", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloWithName("Carlos", "Spanish")
		want := "Hola, Carlos"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := HelloWithName("Carlos", "French")
		want := "Bonjout, Carlos"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}
}
