package stringx

// Clone returns a copy of slice.
//
// notes:
//   - Clone(nil) => nil
//   - Clone([]string{}) => nil
//
// Deprecated: just use `append([]string(nil), original...)` in place.
func Clone(slice []string) []string {
	return append([]string(nil), slice...)
}
