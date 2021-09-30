package types

import (
	"fmt"
)

// Definition of Int64State

var _ CDTState = (*Int64CDTState)(nil)

func NewInt64CDTState() *Int64CDTState {
	return &Int64CDTState{Max: 0, Min: 0}
}

func (s *Int64CDTState) Apply(op OP) error {
	switch op := op.(type) {
	case *AddInt64OP:
		s.Max = op.AddTo(s.Max)
	case *SubInt64OP:
		s.Min = op.AddTo(s.Min)
	default:
		return fmt.Errorf("cannot apply %T", op)
	}
	return nil
}

func (s *Int64CDTState) IsEmpty() bool {
	return s.Max == 0 && s.Min == 0
}

type Int64OPComposer struct{}

var _ OPComposer = (*Int64OPComposer)(nil)

func (Int64OPComposer) Compose(a []OP, b OP) []OP {
	if l := len(a); l == 0 {
		return []OP{b}
	} else if l != 1 {
		panic("fatal error")
	}

	s := a[0]
	return []OP{s.Compose(b)}
}
