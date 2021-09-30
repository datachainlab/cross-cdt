package types

import (
	"github.com/gogo/protobuf/proto"
)

type OP interface {
	proto.Message
	Key() []byte
	Compose(OP) OP
	Inverse() OP
}

type OPManager interface {
	Add(op OP, composer OPComposer)
	OPs() []OP
	GetOPs(key []byte) []OP
}

// NewOPManager returns a OPManager instance
func NewOPManager() *opManager {
	return &opManager{opsm: make(map[string][]OP)}
}

var _ OPManager = (*opManager)(nil)

type opManager struct {
	opsm map[string][]OP
}

func (m *opManager) Add(op OP, composer OPComposer) {
	k := string(op.Key())
	m.opsm[k] = composer.Compose(m.opsm[k], op)
}

func (m opManager) OPs() []OP {
	var ret []OP
	for _, ops := range m.opsm {
		ret = append(ret, ops...)
	}
	return ret
}

func (m opManager) GetOPs(key []byte) []OP {
	k := string(key)
	return m.opsm[k]
}

type OPComposer interface {
	Compose([]OP, OP) []OP
}

type SimpleOPComposer struct{}

var _ OPComposer = (*SimpleOPComposer)(nil)

func (SimpleOPComposer) Compose(a []OP, b OP) []OP {
	c := make([]OP, len(a)+1)
	copy(c[:len(a)], a)
	c[len(a)] = b
	return c
}
