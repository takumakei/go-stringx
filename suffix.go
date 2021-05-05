package stringx

// Suffix treats a string as a suffix string.
type Suffix string

// Map returns a string slice of the result of prefix + s.
func (s Suffix) Map(prefix ...string) []string {
	return s.MapFunc(join, prefix...)
}

// MapPath returns a string slice of the result of filepath.Join(prefix, s).
func (s Suffix) MapPath(prefix ...string) []string {
	return s.MapFunc(pathjoin, prefix...)
}

// MapJoin returns a string slice of the result of
// prefix + sep + s (when both of prefix and s was not empty string),
// s or prefix (of non-empty string).
func (s Suffix) MapJoin(prefix []string, sep string) []string {
	if len(s) == 0 {
		return Clone(prefix)
	}
	mapjoin := func(lhs, rhs string) string {
		if len(lhs) == 0 {
			return rhs
		}
		return lhs + sep + rhs
	}
	return s.MapFunc(mapjoin, prefix...)
}

// MapFunc returns a string slice of the result of fn(prefix, s) for each prefix.
func (s Suffix) MapFunc(fn func(lhs, rhs string) string, prefix ...string) []string {
	a := make([]string, len(prefix))
	for i, v := range prefix {
		a[i] = fn(v, string(s))
	}
	return a
}
