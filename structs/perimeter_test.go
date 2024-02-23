package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, got, want float64, shape Shape) {
		t.Helper()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f", shape, got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		checkPerimeter(t, got, want, rectangle)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := Perimeter(circle)
		want := 62.83185307179586

		checkPerimeter(t, got, want, circle)
	})

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{3, 4, 5}
		got := Perimeter(triangle)
		want := 12.0

		checkPerimeter(t, got, want, triangle)
	})
}

func BenchmarkPerimeter(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		Perimeter(rectangle)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, got, want float64, shape Shape) {
		t.Helper()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f", shape, got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		checkArea(t, got, want, rectangle)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		checkArea(t, got, want, circle)
	})

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{3, 4, 5}
		got := triangle.Area()
		want := 60.0

		checkArea(t, got, want, triangle)
	})
}

func BenchmarkArea(b *testing.B) {
	rectangle := Rectangle{12, 6}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}
