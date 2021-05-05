package stringx_test

import (
	"testing"

	"github.com/takumakei/go-stringx"
)

func TestIndex(t *testing.T) {
	ss := []string{
		"abc", "def", "ghi", "ijk",
		"abc", "def", "ghi", "ijk",
	}
	s := "ijk"

	want := 3
	got := stringx.Index(ss, s)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestLastIndex(t *testing.T) {
	ss := []string{
		"abc", "def", "ghi", "ijk",
		"abc", "def", "ghi", "ijk",
	}
	s := "abc"

	want := 4
	got := stringx.LastIndex(ss, s)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
