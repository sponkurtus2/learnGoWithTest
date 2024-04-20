package structsmethodsinterfaces

import "testing"

// Global var of figures
var (
	rectangle = Rectangle{10.0, 10.0}
	circle    = Circle{10}
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(rectangle)
	var want float64 = 40

	assertMsg(t, got, want)
}

func TestArea(t *testing.T) {
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got: %g, want %g", got, want)
	// 	}
	// }
	//
	// t.Run("Area on rectangle", func(t *testing.T) {
	// 	// got := Area(rectangle)
	// 	// got := rectangle.Area()
	// 	// want := 40.0
	// 	// assertMsg(t, got, want)
	// 	checkArea(t, rectangle, 40.0)
	// })
	//
	// t.Run("Area on circle", func(t *testing.T) {
	// 	// got := Area(Circle)
	// 	// got := circle.Area()
	// 	// want := 314.156
	// 	// assertMsg(t, got, want)
	// 	checkArea(t, circle, 314.1592653589793)
	// })
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Sphere{4.0}, 201.06192982974676},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got: %g, want %g", got, tt.want)
		}
	}
}

func assertMsg(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got: %.2f, want %.2f", got, want)
	}
}
