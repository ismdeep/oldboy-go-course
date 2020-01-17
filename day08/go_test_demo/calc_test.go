package go_test_demo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(2, 4)
	if r != 6 {
		t.Fatal(r)
	}
	t.Logf("test add successfully")
}
