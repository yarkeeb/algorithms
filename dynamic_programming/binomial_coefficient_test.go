package dynamic_programming

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinomial(t *testing.T) {
	res := Binomial(4, 5)
	require.Equal(t, res, (int64)(5), "results not match")
}
