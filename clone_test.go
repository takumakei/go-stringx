package stringx_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takumakei/go-stringx"
)

func TestClone(t *testing.T) {
	tt := []struct {
		A []string
	}{
		{},
		{[]string{}},
		{[]string{"a"}},
		{[]string{"a", "b"}},
	}
	for i, v := range tt {
		got := stringx.Clone(v.A)
		if diff := cmp.Diff(v.A, got); diff != "" {
			t.Errorf("%d -want +got\n%s", i, diff)
		}
	}
}
