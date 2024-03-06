package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("Empty string defaults to 'world'", func(t *testing.T) {
		got := Hello(" ")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Holla, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() //Tells the test suit that this method is a helper
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
