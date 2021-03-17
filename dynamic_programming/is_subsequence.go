package dynamic_programming

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	curr := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[curr] {
			curr++
		}
		if curr == len(s) {
			return true
		}
	}
	return false
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}


func isSubsequenceDynamic(s string, t string) bool {
	n := len(s)
	m := len(t)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if s[i -1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]+1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m] == n
}

