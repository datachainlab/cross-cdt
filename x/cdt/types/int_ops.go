package types

type Int64OP interface {
	OP
	Value() int64
	AddTo(v int64) int64
}

var _ OP = (*AddInt64OP)(nil)

func NewInt64OP(k []byte, v int64) Int64OP {
	if v >= 0 {
		return &AddInt64OP{K: k, V: v}
	} else {
		return &SubInt64OP{K: k, V: -v}
	}
}

func (op AddInt64OP) Key() []byte {
	return op.K
}

func (op AddInt64OP) Value() int64 {
	return op.V
}

func (op AddInt64OP) AddTo(v int64) int64 {
	return v + op.Value()
}

func (op AddInt64OP) Inverse() OP {
	return &AddInt64OP{K: op.K, V: -op.V}
}

func (op AddInt64OP) Compose(other OP) OP {
	return mergeInt64WithAny(&op, other)
}

var _ OP = (*SubInt64OP)(nil)

func (op SubInt64OP) Key() []byte {
	return op.K
}

func (op SubInt64OP) Value() int64 {
	return -op.V
}

func (op SubInt64OP) AddTo(v int64) int64 {
	return v + op.Value()
}

func (op SubInt64OP) Inverse() OP {
	return &SubInt64OP{K: op.K, V: -op.V}
}

func (op SubInt64OP) Compose(other OP) OP {
	return mergeInt64WithAny(&op, other)
}

func mergeInt64WithAny(op Int64OP, any OP) OP {
	switch other := any.(type) {
	case Int64OP:
		return NewInt64OP(other.Key(), other.AddTo(op.Value()))
	default:
		panic("fatal error")
	}
}
