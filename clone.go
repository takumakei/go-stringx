package stringx

// Clone returns a copy of slice.
func Clone(slice []string) []string {
	if slice == nil {
		return nil
	}
	a := make([]string, len(slice))
	copy(a, slice)
	return a
}
