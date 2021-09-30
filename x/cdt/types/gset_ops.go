package types

import "bytes"

var _ OP = (*GSetOP)(nil)

func NewGSetOP(k []byte, add [][]byte) *GSetOP {
	return &GSetOP{K: k, Add: add}
}

func (op GSetOP) Key() []byte {
	return op.K
}

func (op GSetOP) Compose(other OP) OP {
	if !bytes.Equal(op.K, other.Key()) {
		panic("key mismatch")
	}
	switch other := other.(type) {
	case *GSetOP:
		return NewGSetOP(op.Key(), op.AddTo(other.Add))
	default:
		panic("fatal error")
	}
}

func (op GSetOP) Inverse() OP {
	inv := GSetInverseOP(op)
	return &inv
}

func (op GSetOP) AddTo(elements [][]byte) [][]byte {
	return append(op.Add, elements...)
}

var _ OP = (*GSetInverseOP)(nil)

func NewGSetInverseOP(k []byte, add [][]byte) *GSetInverseOP {
	return &GSetInverseOP{K: k, Add: add}
}

func (op GSetInverseOP) Key() []byte {
	return op.K
}

func (op GSetInverseOP) Compose(other OP) OP {
	if !bytes.Equal(op.K, other.Key()) {
		panic("key mismatch")
	}
	switch other := other.(type) {
	case *GSetInverseOP:
		return NewGSetInverseOP(other.Key(), other.AddTo(op.Add))
	default:
		panic("fatal error")
	}
}

func (op GSetInverseOP) Inverse() OP {
	panic("not implemented error")
}

func (op GSetInverseOP) AddTo(elements [][]byte) [][]byte {
	return append(op.Add, elements...)
}
