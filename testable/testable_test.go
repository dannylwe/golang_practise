package testable

import "testing"

func TestStrInSlice(t *testing.T){
	got := StrInSlice([]string{"a", "b"}, "c")
	if got != false {
		t.Errorf("expecting false got true")
	}
}