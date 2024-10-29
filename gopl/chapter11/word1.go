package chapter11

// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome1(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
