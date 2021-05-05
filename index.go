package stringx

// Index returns the index of the first element of target in ss, or -1 if
// target is not present in ss.
func Index(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last element of target in ss, or -1 if
// target is not present in ss.
func LastIndex(slice []string, target string) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
