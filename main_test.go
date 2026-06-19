package nemo

import (
	"testing"
)

func TestLatinToUtf8(t *testing.T) {
	source := []byte("hoge\x10uhauha\x81\x82")
	expect := `hoge\x10uhauha\x81\x82`
	result := latinToUtf8(source)
	if expect != result {
		t.Fatalf("expect %q, but got %q", expect, result)
	}
}
