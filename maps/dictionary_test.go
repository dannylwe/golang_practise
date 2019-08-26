package dict

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test"}
	
	// search by key
	got := Search(dictionary, "test")
	want := "this is a test"
	assertStrings(t, got, want)

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted '%s' but got '%s' given '%s'", want, got, "test")
	}
}
