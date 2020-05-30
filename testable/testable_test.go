package testable

import "testing"

func TestStrInSlice(t *testing.T) {

	var tests = []struct {
		slice []string
		find  string
		want  bool
	}{
		{[]string{"a", "b"}, "c", false},
		{[]string{"a", "b"}, "b", true},
	}

	// got := StrInSlice([]string{"a", "b"}, "c")
	// if got != false {
	// 	t.Errorf("expecting false got true")
	// }

	for _, test := range tests {
		t.Run("table tests for slice", func(t *testing.T) {
			got := StrInSlice(test.slice, test.find)
			if got != test.want {
				t.Errorf("expecting %v got %v", test.want, test.find)
			}
		})
	}
}
