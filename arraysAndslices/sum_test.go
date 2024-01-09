package main

import "testing"

func TestSum(t *testing.T) {
	// t.Run("Collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}
	//
	// 	got := Sum(numbers)
	// 	want := 15
	//
	// 	assertCorrectMessage(t, got, want, numbers)
	// })

	t.Run("Collection of any size numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
	})
}

func assertCorrectMessage(t testing.TB, got int, want int, numbers []int) {
	// Tells the test suite that the method is a helper.
	// Will provide line number inside of failed function on fail
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
