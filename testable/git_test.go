package testable

import (
	"testing"
)

func TestGetAverageStarsByUsername(t *testing.T) {

	var tests = []struct {
		username string
		want     float64
	}{
		{"ocotocat", 9.5},
		
	}

	// got := StrInSlice([]string{"a", "b"}, "c")
	// if got != false {
	// 	t.Errorf("expecting false got true")
	// }
	
	mock := new(Mock)
	for _, test := range tests {
		t.Run("Get Average Stars Github", func(t *testing.T){
			got, err:= GetAverageStarsByUsername(mock, test.username)
			if err != nil {
				t.Errorf("expecting nil err got %v", err)

			}
			if got != test.want {
				t.Errorf("expecting %v got %v", test.want, got)
			}
		})
	}
}
