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

func TestAreaRefactor(t *testing.T) {
	areaTests := []struct {
		shape Shape
		hasArea float64
		name string
	}{
		{shape: Rectangle{Width: 12, Length: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape:Triangle{Base:12, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T){
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f hasArea %.2f", tt.shape, got, tt.hasArea)
			}	
		})
		
	}
}
