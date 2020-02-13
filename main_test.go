package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	text := hello()
	expect := "hello, GitHub Actions"
	if text != expect {
		t.Errorf("expect %s, got %s", expect, text)
	}
}
