package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGSetValue(t *testing.T) {
	require := require.New(t)
	set := newSetValue(1, 2)
	require.Len(set.Elements, 2)
	set.Add([]byte(fmt.Sprint(1)))
	require.Len(set.Elements, 2)
	set.Add([]byte(fmt.Sprint(2)))
	require.Len(set.Elements, 2)
	set.Add([]byte(fmt.Sprint(3)))
	require.Len(set.Elements, 3)
	set.Add([]byte(fmt.Sprint(3)), []byte(fmt.Sprint(4)))
	require.Len(set.Elements, 4)
}

func newSetValue(vs ...int) *GSetValue {
	var elements []Element
	for _, v := range vs {
		elements = append(elements, []byte(fmt.Sprint(v)))
	}
	set := NewGSetValue(elements...)
	return set
}
