package githubactionstestinggo

import (
	"testing"
)

func TestSum(t *testing.T) {
	a := 10
	b := 15

	e := 25

	r := Sum(a, b)

	if r != e {
		t.Fatalf("Incorrect result. Expected %d, got %d", e, r)
	}
}
