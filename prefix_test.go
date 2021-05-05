package stringx_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takumakei/go-stringx"
)

func TestPrefix_MapJoin(t *testing.T) {
	want := []string{"a", "b"}
	got := stringx.Prefix("").MapJoin(want, ":")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("-want +got\n%s", diff)
	}
}

func ExamplePrefix_Map() {
	for _, v := range stringx.Prefix("hello ").Map("alice", "bob") {
		fmt.Println(v)
	}
	// Output:
	// hello alice
	// hello bob
}

func ExamplePrefix_MapPath() {
	for _, v := range stringx.Prefix("/home/user").MapPath(".config", ".local") {
		fmt.Println(v)
	}
	// Output:
	// /home/user/.config
	// /home/user/.local
}

func ExamplePrefix_MapJoin() {
	for _, v := range stringx.Prefix("x").MapJoin([]string{"txt", ""}, ".") {
		fmt.Println(v)
	}
	// Output:
	// x.txt
	// x
}
