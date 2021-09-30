package keeper

import (
	"fmt"

	"github.com/datachainlab/cross-cdt/x/cdt/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

// CDTStore manages cdt states of each key
type CDTStore struct {
	m      codec.Codec
	kvs    types.KVStoreI
	schema types.Schema
}

func NewCDTStore(m codec.Codec, kvs types.KVStoreI, schema types.Schema) CDTStore {
	return CDTStore{m: m, kvs: kvs, schema: schema}
}

// ApplyOP applies the op into the state
func (s CDTStore) ApplyOP(ctx sdk.Context, op types.OP) error {
	state := s.GetStateOrEmpty(ctx, op.Key())
	if err := state.Apply(op); err != nil {
		return err
	}
	s.setState(ctx, op.Key(), state)
	return nil
}

// GetState returns the state corresponding to the key
func (s CDTStore) GetState(ctx sdk.Context, key []byte) types.CDTState {
	bz := s.kvs.Get(ctx, key)
	if bz == nil {
		return nil
	}
	var anyState types.AnyCDTState
	if err := proto.Unmarshal(bz, &anyState); err != nil {
		panic(err)
	}
	state, err := types.UnpackCDTState(s.m, &anyState)
	if err != nil {
		panic(err)
	}
	return state
}

// HasState returns a boolean value whether if the state corresponding to the key
func (s CDTStore) HasState(ctx sdk.Context, key []byte) bool {
	return s.kvs.Get(ctx, key) != nil
}

// GetStateOrEmpty returns the state corresponding to the key if it exists, otherwise it returns an empty state
func (s CDTStore) GetStateOrEmpty(ctx sdk.Context, key []byte) types.CDTState {
	current := s.GetState(ctx, key)
	if current == nil {
		ct, ok := s.schema.MatchType(key)
		if !ok {
			panic(fmt.Errorf("schema not found: %v", key))
		}
		// use default state
		current = types.GetEmptyCDTState(ct)
	}
	return current
}

func (s CDTStore) setState(ctx sdk.Context, key []byte, state types.CDTState) {
	if state.IsEmpty() {
		s.kvs.Delete(ctx, key)
		return
	}
	bz, err := proto.Marshal(types.PackCDTState(state))
	if err != nil {
		panic(err)
	}
	s.kvs.Set(ctx, key, bz)
}
