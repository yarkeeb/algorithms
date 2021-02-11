package dynamic_programming

func max(n, m int64) int64 {
	if n < m {
		return m
	}
	return n
}

func Binomial(n, m int64) int64 {
	var i, j int64
	d := max(n, m)
	c := make([][]int64, d+1)
	for i := range c {
		c[i] = make([]int64, d+1)
	}
	for i = 0; i <= d; i++ {
		c[i][0] = 1
	}
	for j = 0; j <= d; j++ {
		c[j][j] = 1
	}
	for i = 1; i <= d; i++ {
		for j = 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	return c[m][n]
}