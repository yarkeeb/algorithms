package dynamic_programming

//Leetcode 746
//https://leetcode.com/problems/min-cost-climbing-stairs/
//min cost climbing stairs
//The core approach here is to look at the problem from "end"
//What is the cost of climbing stairs from the last step? And from the previous?

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	c1 := cost[size - 1]
	c2 := cost[size - 2]
	for i := size - 3; i >= 0; i-- {
		curr := cost[i] + min(c1, c2)
		c1 = c2
		c2 = curr
	}
	return min(c1, c2)
}