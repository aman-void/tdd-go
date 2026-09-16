package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ken Thompson!", "")
		want := "Hello, Ken Thompson!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Rob Pike", "Spanish")
		want := "Hola, Rob Pike"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Robert Griesemer", "French")
		want := "Bonjour, Robert Griesemer"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Discipline
// - Write a test
// - Make the compiler pass
// - Run the test, see that it fails and check the error message is meaningful
// - Write enough code to make the test pass
// - Refactor
