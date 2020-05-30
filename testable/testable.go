package testable

func StrInSlice(slice []string, find string) bool {
	for _, v := range slice {
		if v == find {
			return true
		}
	}
	return false
}
