package pager

import (
	"testing"
)

func TestTrimLeft(t *testing.T) {
	src := "\x1B[31m123456789"
	result := trimLeft(src, 5)
	expect := "\x1B[31m6789"
	if result != expect {
		t.Fatalf("expect %q, but got %q", expect, result)
	}
}

func TestTruncate(t *testing.T) {
	src := "\x1B[31m123456789"
	result := Truncate(src, 5)
	expect := "\x1b[31m1234"
	if result != expect {
		t.Fatalf("expect %q, but got %q", expect, result)
	}
}

func TestBoth(t *testing.T) {
	src := "\x1B[31m123456789"
	result := Truncate(trimLeft(src, 3), 5)
	expect := "\x1B[31m4567"
	if result != expect {
		t.Fatalf("expect %q, but got %q", expect, result)
	}
}
