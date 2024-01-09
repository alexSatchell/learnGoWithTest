package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("Iterating 'a' 5 times", func(t *testing.T) {
		got := Iterate("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Iterating 'a' for specifed amount", func(t *testing.T) {
		got := Iterate("a", 3)
		want := "aaa"

		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 5)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Tells the test suite that the method is a helper.
	// Will provide line number inside of failed function on fail
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
