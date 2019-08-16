package shapes

import "testing"
func TestPerimeter(t *testing.T){
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f but wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T){
	got := Area(9.0, 10.0)
	want := 90.0
	if got != want {
		t.Errorf("got %.2f but wanted %.2f", got, want)
	}
}
