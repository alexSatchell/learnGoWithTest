package search

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Search known word", func(t *testing.T) {
		definition, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, definition, want)
	})

	t.Run("Search unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		asserError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	dictionary.Add(word, definition)
	asserDefinition(t, dictionary, word, definition)
}

// helpers
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func asserError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func asserDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word")
	}
	assertStrings(t, got, definition)
}
