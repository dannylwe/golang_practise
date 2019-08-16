package shapes

import "testing"
func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}
func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    }
	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{9.0, 10.0}
		want := 90.0
		checkArea(t, rectangle, want)
	})
	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
	
}
