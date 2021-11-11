package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOPManager(t *testing.T) {
	require := require.New(t)
	expected := []OP{
		NewInt64OP([]byte("a"), 1),
		NewInt64OP([]byte("b"), 2),
		NewInt64OP([]byte("c"), 3),
		NewInt64OP([]byte("d"), 4),
		NewInt64OP([]byte("e"), 5),
	}

	opManager := NewOPManager()
	composer := SimpleOPComposer{}
	opManager.Add(expected[0], composer)
	opManager.Add(expected[4], composer)
	opManager.Add(expected[2], composer)
	opManager.Add(expected[3], composer)
	opManager.Add(expected[1], composer)
	ops := opManager.OPs()
	require.Len(ops, 5)
	require.Equal(expected, ops)
}
