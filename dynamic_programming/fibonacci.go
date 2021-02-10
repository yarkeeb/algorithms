package dynamic_programming

func FibonacciRecursive(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciRecursive(n - 1) + FibonacciRecursive(n - 2)
}

func FibonacciDynamic(n int64) int64 {
	hash := make(map[int64]int64)
	hash[0] = 0
	hash[1] = 1
	for i := int64(2); i <= n; i++ {
		if _, ok := hash[i]; !ok {
			hash[i] = hash[i - 1] + hash[i - 2]
		}
	}
	return hash[n]
}

func FibonacciDynamicNoExtraMemory(n int64) int64 {
	var i, prev, next int64 = 0, 0, 1
	for i = int64(2); i <= n; i++ {
		curr := prev + next
		prev = next
		next = curr
	}
	return next
}