package linearsearch

import "testing"

func TestLinearSearchOrdered(t *testing.T) {
	orderedArr := []int{1, 3, 5, 7, 9}

	got := LinearSearchOrdered(orderedArr, 3)
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
