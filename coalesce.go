package stringx

// Coalesce returns the first non-empty string of s.
func Coalesce(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	last := len(s) - 1
	for _, v := range s[:last] {
		if len(v) > 0 {
			return v
		}
	}
	return s[last]
}
