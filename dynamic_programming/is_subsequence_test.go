package dynamic_programming

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	res := isSubsequence("abc", "ahbsdfcd")
	require.True(t, res,"results not match")
	res = isSubsequence("klm", "abvsde")
	require.False(t, res, "results not match")
	res = isSubsequenceDynamic("abc", "ahbsdfcd")
	require.True(t, res,"results not match")
	res = isSubsequenceDynamic("klm", "abvsde")
	require.False(t, res, "results not match")
}


