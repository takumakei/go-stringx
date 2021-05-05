package stringx

import (
	"path/filepath"
)

func join(lhs, rhs string) string {
	return lhs + rhs
}

func pathjoin(lhs, rhs string) string {
	return filepath.Join(lhs, rhs)
}
