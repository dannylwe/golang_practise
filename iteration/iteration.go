package iteration

const repeatCount = 3

// Repeat returns string of repeated chars
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
	}
	return repeated
}
