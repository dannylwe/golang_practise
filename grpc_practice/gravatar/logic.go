package gravatar

import (
	"crypto/md5"
	"fmt"
)

func generateHash(email string) [16]byte {
	return md5.Sum([]byte(email))
}

func gravatarURL(hash [16]byte, size uint32) string {
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d", hash, size)
}

// Gravatar function generates url with md5 and size
func Gravatar(email string, size uint32) string {
	hash := generateHash(email)
	return gravatarURL(hash, size)
}