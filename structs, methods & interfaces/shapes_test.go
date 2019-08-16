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
	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{9.0, 10.0}
		got := rectangle.Area()
		want := 90.0
		if got != want {
			t.Errorf("got %.2f but wanted %.2f", got, want)
		}	
	})
	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %.2f but wanted %.2f", got, want)
		}	
	})
	
}
