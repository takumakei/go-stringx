package stringx_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takumakei/go-stringx"
)

func TestSuffix_MapJoin(t *testing.T) {
	want := []string{"a", "b"}
	got := stringx.Suffix("").MapJoin(want, ":")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("-want +got\n%s", diff)
	}
}

func ExampleSuffix_Map() {
	for _, v := range stringx.Suffix(".txt").Map("alice", "bob") {
		fmt.Println(v)
	}
	// Output:
	// alice.txt
	// bob.txt
}

func ExampleSuffix_MapPath() {
	for _, v := range stringx.Suffix("bin").MapPath("/usr/local", "/usr") {
		fmt.Println(v)
	}
	// Output:
	// /usr/local/bin
	// /usr/bin
}

func ExampleSuffix_MapJoin() {
	for _, v := range stringx.Suffix("Func()").MapJoin([]string{"A", ""}, ".") {
		fmt.Println(v)
	}
	// Output:
	// A.Func()
	// Func()
}
