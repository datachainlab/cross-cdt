package keeper

import (
	"fmt"

	"github.com/datachainlab/cross-cdt/x/cdt/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	contracttypes "github.com/datachainlab/cross/x/core/contract/types"
)

type GSetStore struct {
	cdc    codec.Codec
	prefix []byte

	schema     types.Schema
	stateStore types.KVStoreI
	cdtStore   CDTStore
}

func NewGSetStore(
	cdc codec.Codec,
	schema types.Schema,
	stateStore types.KVStoreI,
	cdtStore CDTStore,
	prefix []byte,
) GSetStore {
	return GSetStore{
		cdc:        cdc,
		prefix:     prefix,
		schema:     schema,
		stateStore: stateStore,
		cdtStore:   cdtStore,
	}
}

func (s GSetStore) Get(ctx sdk.Context, key []byte) []types.Element {
	fullKey := buildKey(s.prefix, key)
	s.schema.MustMatchGSet(fullKey)

	set := s.getCommittedState(ctx, fullKey)
	switch contracttypes.CommitModeFromContext(ctx.Context()) {
	case contracttypes.UnspecifiedMode, contracttypes.BasicMode:
	case contracttypes.AtomicMode:
		ops := types.OPManagerFromContext(ctx.Context()).GetOPs(fullKey)
		if len(ops) == 1 {
			switch op := ops[0].(type) {
			case *types.GSetOP:
				set.Add(op.Add...)
			default:
				panic(fmt.Errorf("unknown type %T", op))
			}
		}
	default:
		panic(fmt.Sprintf("unknown mode '%v'", contracttypes.CommitModeFromContext(ctx.Context())))
	}
	if set == nil {
		return nil
	}
	return set.Elements
}

func (s GSetStore) Add(ctx sdk.Context, key []byte, elements ...types.Element) {
	fullKey := buildKey(s.prefix, key)
	s.schema.MustMatchGSet(fullKey)

	switch contracttypes.CommitModeFromContext(ctx.Context()) {
	case contracttypes.UnspecifiedMode, contracttypes.BasicMode:
		s.addCommittedState(ctx, fullKey, elements...)
	case contracttypes.AtomicMode:
		var els []types.Element
		committedState := s.getCommittedState(ctx, key)
		for _, el := range elements {
			if !committedState.Lookup([]types.Element{el}) {
				els = append(els, el)
			}
		}
		types.OPManagerFromContext(ctx.Context()).Add(
			types.NewGSetOP(fullKey, els),
			types.GSetOPComposer{},
		)
	}
}

func (s GSetStore) Lookup(ctx sdk.Context, key []byte, elements ...types.Element) bool {
	fullKey := buildKey(s.prefix, key)
	s.schema.MustMatchGSet(fullKey)

	cdtState := s.cdtStore.GetStateOrEmpty(ctx, fullKey).(*types.GSetCDTState)
	committedState := s.getCommittedState(ctx, fullKey)

	switch contracttypes.CommitModeFromContext(ctx.Context()) {
	case contracttypes.AtomicMode:
		ops := types.OPManagerFromContext(ctx.Context()).GetOPs(fullKey)
		if l := len(ops); l == 1 {
			switch op := ops[0].(type) {
			case *types.GSetOP:
				set := types.NewGSetValue(op.Add...)
				if set.Lookup(elements) {
					return true
				}
			default:
				panic(fmt.Errorf("unknown type %T", op))
			}
		} else if l > 1 {
			panic("fatal error")
		}
	}

	if committedState.Lookup(elements) {
		return true
	} else if !cdtState.Lookup(elements) {
		return false
	} else {
		panic(types.ErrIndefiniteState)
	}
}

func (s GSetStore) setCommittedState(ctx sdk.Context, key []byte, set *types.GSetValue) {
	bz, err := s.cdc.Marshal(set)
	if err != nil {
		panic(err)
	}
	s.stateStore.Set(ctx, key, bz)
}

func (s GSetStore) getCommittedState(ctx sdk.Context, key []byte) *types.GSetValue {
	var set types.GSetValue
	v := s.stateStore.Get(ctx, key)
	if v != nil {
		if err := s.cdc.Unmarshal(v, &set); err != nil {
			panic(err)
		}
	}
	return &set
}

func (s GSetStore) addCommittedState(ctx sdk.Context, key []byte, elements ...types.Element) {
	set := s.getCommittedState(ctx, key)
	set.Add(elements...)
	s.setCommittedState(ctx, key, set)
}
