package dynamic_programming


import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	res := maxProfit([]int{7,1,5,3,6,4})
	require.Equal(t, res, 5, "results not match")
	res = maxProfit([]int{7,6,4,3,1})
	require.Equal(t, res, 0, "results not match")
}

