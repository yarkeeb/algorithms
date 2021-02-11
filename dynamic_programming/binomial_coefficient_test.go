package dynamic_programming

import "testing"

func TestBinomial(t *testing.T) {
	res := Binomial(4, 5)
	if res != 5 {
		t.Errorf("Binomial was incorrect, got: %d, expected: %d", res, 5)
	}

}
