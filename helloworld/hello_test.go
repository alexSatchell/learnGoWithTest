package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Alex", "")
		want := "Hello, Alex!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sebastian", "French")
		want := "Bonjour, Sebastian!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Tells the test suite that the method is a helper.
	// Will provide line number inside of failed function on fail
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
