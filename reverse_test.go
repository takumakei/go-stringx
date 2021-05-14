package stringx_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takumakei/go-stringx"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		In   []string
		Want []string
	}{
		{nil, nil},
		{[]string{}, []string{}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
	}
	for i, v := range cases {
		stringx.Reverse(v.In)
		if diff := cmp.Diff(v.Want, v.In); len(diff) > 0 {
			t.Errorf("[%d] -want +got\n%s", i, diff)
		}
	}
}

func ExampleReverse() {
	slice := []string{
		"grape",
		"lemon",
		"melon",
		"orange",
	}
	stringx.Reverse(slice)
	fmt.Println(strings.Join(slice, ","))
	// Output:
	// orange,melon,lemon,grape
}
