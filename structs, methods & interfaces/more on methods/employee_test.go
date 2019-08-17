package employee

import "testing"

func TestFullName(t *testing.T) {
	employee := Employee{"Daniel", "Roy"}
	want := "Daniel Roy"
	got := employee.FullName()
	if got != want {
		t.Errorf("got %v but expected %v", got, want)
	}
}

func TestChangeName(t *testing.T) {
	employee := Employee{"Daniel", "Roy"}
	updateFirstName := "Kanye west"
	want := updateFirstName + " " + "Roy"
	employee.changeNameFirstName(updateFirstName)
	got := employee.FullName()
	if got != want {
		t.Errorf("got %v but expected %v", got, want)
	}
}
