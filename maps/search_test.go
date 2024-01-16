package search

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	// t.Run("Search known word", func(t *testing.T) {
	// 	_, got := dictionary.Search("test")
	// 	want := "this is just a test"
	//
	// 	assertStrings(t, got, want)
	// })

	t.Run("Search unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertStrings(t, got, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
