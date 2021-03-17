package dynamic_programming

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	res := ClimbStairs(6)
	require.Equal(t, res, 6, "results not equal")
	res = ClimbStairs(1)
	require.Equal(t, res, 1, "results not equal")
	res = ClimbStairs(2)
	require.Equal(t, res, 2, "results not equal")
}

