package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // marks function as test helper
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, world' when an empty string is applied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Steve", "French")
		want := "Bonjour, Steve"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Te Reo", func(t *testing.T) {
		got := Hello("Ben", "Te Reo")
		want := "Kia Ora, Ben"
		assertCorrectMessage(t, got, want)
	})

}
