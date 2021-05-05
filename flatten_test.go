package stringx_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takumakei/go-stringx"
)

func TestFlatten(t *testing.T) {
	want := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	got := stringx.Flatten(
		"a",
		[]string{"b", "c"},
		[...]string{"d", "e"},
		[][]string{{"f"}, {"g", "h"}},
	)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("-want +got\n%s", diff)
	}
}
