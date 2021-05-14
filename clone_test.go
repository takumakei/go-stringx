package stringx_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takumakei/go-stringx"
)

func TestClone(t *testing.T) {
	cases := []struct {
		In   []string
		Want []string
	}{
		{},
		{[]string{}, nil},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
	}
	for i, c := range cases {
		got := stringx.Clone(c.In)
		if diff := cmp.Diff(c.Want, got); diff != "" {
			t.Errorf("%d -want +got\n%s", i, diff)
		}
	}
}
