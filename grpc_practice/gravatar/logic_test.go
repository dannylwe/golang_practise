package gravatar

import (
	"crypto/md5"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGravatar(t *testing.T) {
	t.Run("generate url", func(t *testing.T) {

		var size uint32 = 10
		endpoint := "https://www.gravatar.com/avatar/46040c38d1cbe8ffcd3df6c8ba787951?s=%v"
		email := "abcd@gmail.com"
	
		generatedURL := gravatar(email, size)
		expected := fmt.Sprintf(endpoint, size)
	
		if expected != generatedURL{
			t.Errorf("expected %v but got %v", expected, generatedURL)
		}
	})
	t.Run("generate Hash", func(t *testing.T) {
		email := "xxjdjj@gmail.com"
		got := generateHash(email)
		expected := md5.Sum([]byte("nfijrnkerj"))

		assert.NotEqual(t, got, expected)
	})
}
