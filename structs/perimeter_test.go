package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{6.0, 4.0}

	got := Perimeter(rectangle)
	want := 20.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
//
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
//
// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{6.0, 4.0}
// 		want := 24.0
// 		checkArea(t, rectangle, want)
// 	})
//
// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }

// Example of table driven test
func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Height: 6.0, Width: 4.0}, want: 24.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
