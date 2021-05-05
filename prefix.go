package stringx

// Prefix treats a string as a prefix string.
type Prefix string

// Map returns a string slice of the result of s + suffix.
func (s Prefix) Map(suffix ...string) []string {
	return s.MapFunc(join, suffix...)
}

// MapPath returns a string slice of the result of filepath.Join(s, suffix).
func (s Prefix) MapPath(suffix ...string) []string {
	return s.MapFunc(pathjoin, suffix...)
}

// MapJoin returns a string slice of the result of
// s + sep + suffix (when both of s and suffix was not empty string),
// s or suffix (of non-empty string).
func (s Prefix) MapJoin(suffix []string, sep string) []string {
	if len(s) == 0 {
		return Clone(suffix)
	}
	mapjoin := func(lhs, rhs string) string {
		if len(rhs) == 0 {
			return lhs
		}
		return lhs + sep + rhs
	}
	return s.MapFunc(mapjoin, suffix...)
}

// MapFunc returns a string slice of the result of fn(s, suffix) for each suffix.
func (s Prefix) MapFunc(fn func(lhs, rhs string) string, suffix ...string) []string {
	a := make([]string, len(suffix))
	for i, v := range suffix {
		a[i] = fn(string(s), v)
	}
	return a
}
