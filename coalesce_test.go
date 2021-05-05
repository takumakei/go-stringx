package stringx_test

import (
	"testing"

	"github.com/takumakei/go-stringx"
)

func TestCoalesce(t *testing.T) {
	tt := []struct {
		A []string
		W string
	}{
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{"", ""}, ""},
		{[]string{"a"}, "a"},
		{[]string{"", "a"}, "a"},
		{[]string{"", "", "a"}, "a"},
		{[]string{"a", "b"}, "a"},
		{[]string{"", "a", "b"}, "a"},
		{[]string{"", "", "a", "b"}, "a"},
	}
	for i, v := range tt {
		got := stringx.Coalesce(v.A...)
		if got != v.W {
			t.Errorf("%d got %v, want %v", i, got, v.W)
		}
	}
}
