package keeper

import (
	"fmt"

	"github.com/datachainlab/cross-cdt/x/cdt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	contracttypes "github.com/datachainlab/cross/x/core/contract/types"
)

type Int64Store struct {
	prefix []byte

	schema     types.Schema
	stateStore types.KVStoreI
	cdtStore   CDTStore
}

func NewInt64Store(
	schema types.Schema,
	stateStore types.KVStoreI,
	cdtStore CDTStore,
	prefix []byte,
) Int64Store {
	return Int64Store{
		prefix:     prefix,
		schema:     schema,
		stateStore: stateStore,
		cdtStore:   cdtStore,
	}
}

func (s Int64Store) Get(ctx sdk.Context, key []byte) int64 {
	fullKey := buildKey(s.prefix, key)
	s.schema.MustMatchInt64(fullKey)

	switch contracttypes.CommitModeFromContext(ctx.Context()) {
	case contracttypes.UnspecifiedMode, contracttypes.BasicMode:
		return s.stateStore.GetInt64(ctx, fullKey)
	case contracttypes.AtomicMode:
		ops := types.OPManagerFromContext(ctx.Context()).GetOPs(fullKey)
		if len(ops) == 1 {
			switch op := ops[0].(type) {
			case types.Int64OP:
				return op.AddTo(s.stateStore.GetInt64(ctx, fullKey))
			default:
				panic(fmt.Errorf("unknown type %T", op))
			}
		} else {
			return s.stateStore.GetInt64(ctx, fullKey)
		}
	default:
		panic(fmt.Errorf("unknown type"))
	}
}

func (s Int64Store) Add(ctx sdk.Context, key []byte, value int64) {
	fullKey := buildKey(s.prefix, key)
	s.schema.MustMatchInt64(fullKey)

	switch contracttypes.CommitModeFromContext(ctx.Context()) {
	case contracttypes.UnspecifiedMode, contracttypes.BasicMode:
		v := s.stateStore.GetInt64(ctx, fullKey)
		s.stateStore.SetInt64(ctx, fullKey, v+value)
	case contracttypes.AtomicMode:
		types.OPManagerFromContext(ctx.Context()).Add(
			types.NewInt64OP(fullKey, value),
			types.Int64OPComposer{},
		)
	}
}

func (s Int64Store) GTE(ctx sdk.Context, key []byte, value int64) bool {
	fullKey := buildKey(s.prefix, key)
	s.schema.MustMatchInt64(fullKey)

	cdtState := s.cdtStore.GetStateOrEmpty(ctx, fullKey).(*types.Int64CDTState)
	v := s.stateStore.GetInt64(ctx, fullKey)

	if contracttypes.CommitModeFromContext(ctx.Context()) == contracttypes.AtomicMode {
		ops := types.OPManagerFromContext(ctx.Context()).GetOPs(fullKey)
		if len(ops) == 1 {
			switch op := ops[0].(type) {
			case types.Int64OP:
				v = op.AddTo(v)
			default:
				panic(fmt.Errorf("unknown type %T", op))
			}
		} else {
		}
	}

	if v+cdtState.Min >= value {
		return true
	} else if v+cdtState.Max < value {
		return false
	} else {
		panic(types.ErrIndefiniteState)
	}
}

func (s Int64Store) GT(ctx sdk.Context, key []byte, value int64) bool {
	fullKey := buildKey(s.prefix, key)
	s.schema.MustMatchInt64(fullKey)

	cdtState := s.cdtStore.GetStateOrEmpty(ctx, fullKey).(*types.Int64CDTState)
	v := s.stateStore.GetInt64(ctx, fullKey)

	if contracttypes.CommitModeFromContext(ctx.Context()) == contracttypes.AtomicMode {
		ops := types.OPManagerFromContext(ctx.Context()).GetOPs(fullKey)
		if len(ops) == 1 {
			switch op := ops[0].(type) {
			case types.Int64OP:
				v = op.AddTo(v)
			default:
				panic(fmt.Errorf("unknown type %T", op))
			}
		} else {
		}
	}

	if v+cdtState.Min > value {
		return true
	} else if v+cdtState.Max <= value {
		return false
	} else {
		panic(types.ErrIndefiniteState)
	}
}

func (s Int64Store) LTE(ctx sdk.Context, key []byte, value int64) bool {
	return !s.GT(ctx, key, value)
}

func (s Int64Store) LT(ctx sdk.Context, key []byte, value int64) bool {
	return !s.GTE(ctx, key, value)
}
