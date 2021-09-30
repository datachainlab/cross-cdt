package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterInterfaces register the ibc transfer module interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*OP)(nil),
		&AddInt64OP{},
		&SubInt64OP{},
		&GSetOP{},
	)
	registry.RegisterImplementations(
		(*CDTState)(nil),
		&Int64CDTState{},
		&GSetCDTState{},
	)
}

var (
	// ModuleCdc references the global x/ibc-transfer module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to x/ibc-transfer and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)

func PackOPs(ops []OP) *AnyOPs {
	var anys []codectypes.Any
	for _, op := range ops {
		any := codectypes.UnsafePackAny(op)
		anys = append(anys, *any)
	}
	return &AnyOPs{Ops: anys}
}

func UnpackOPs(m codec.Codec, anyOPs []codectypes.Any) ([]OP, error) {
	ops := make([]OP, len(anyOPs))
	for i, anyOP := range anyOPs {
		var op OP
		if err := m.UnpackAny(&anyOP, &op); err != nil {
			return nil, err
		}
		ops[i] = op
	}
	return ops, nil
}

func PackCDTState(s CDTState) *AnyCDTState {
	any := codectypes.UnsafePackAny(s)
	return &AnyCDTState{State: *any}
}

func UnpackCDTState(m codec.Codec, any *AnyCDTState) (CDTState, error) {
	var state CDTState
	if err := m.UnpackAny(&any.State, &state); err != nil {
		return nil, err
	}
	return state, nil
}
