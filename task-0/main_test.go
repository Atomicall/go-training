package main

import (
	h "hello/hello"
	"testing"
)

func testHello(t *testing.T) {
	expected := "Hello, world"
	res := h.Hello()
	if expected != res {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res, expected)
	}
}
