package dynamic_programming

import "testing"

func TestMaxSubArray(t *testing.T) {
	t1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if res := MaxSubArray(t1); res != 6 {
		t.Errorf("MaxSubArray result incorrect, got: %d, expected: %d", res, 6)
	}
	t2 := []int{1}
	if res := MaxSubArray(t2); res != 1 {
		t.Errorf("MaxSubArray result incorrect, got: %d, expected: %d", res, 1)
	}
	t3 := []int{-10000}
	if res := MaxSubArray(t3); res != -10000 {
		t.Errorf("MaxSubArray result incorrect, got: %d, expected: %d", res, -10000)
	}
}