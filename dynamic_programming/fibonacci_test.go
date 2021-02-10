package dynamic_programming

import (
	"testing"
)

func TestFibonacciRecursive(t *testing.T) {
	res := FibonacciRecursive(10)
	if res != 55 {
		t.Errorf("Fibonacci recursive was incorrect, got: %d, expected: %d", res, 55)
	}
}
func TestFibonacciDynamic(t *testing.T) {
	res := FibonacciDynamic(10)
	if res != 55 {
		t.Errorf("Fibonacci dynamic was incorrect, got: %d, expected: %d", res, 55)
	}
}

func TestFibonacciDynamicNoExtraMemory(t *testing.T) {
	res := FibonacciDynamicNoExtraMemory(10)
	if res != 55 {
		t.Errorf("Fibonacci dynamic was incorrect, got: %d, expected: %d", res, 55)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciRecursive(20)
	}
}

func BenchmarkFibonacciDynamic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciDynamic(20)
	}
}

func BenchmarkFibonacciDynamicNoExtraMemory(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciDynamicNoExtraMemory(20)
	}
}
