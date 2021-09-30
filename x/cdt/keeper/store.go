package keeper

import (
	"fmt"

	"github.com/datachainlab/cross-cdt/x/cdt/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	contracttypes "github.com/datachainlab/cross/x/core/contract/types"
	"github.com/gogo/protobuf/proto"
)

var _ contracttypes.CommitStoreI = (*Store)(nil)

type Store struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey

	schema     types.Schema
	stateStore types.KVStoreI
	txStore    types.KVStoreI
	cdtStore   CDTStore
}

func NewStore(cdc codec.Codec, storeKey sdk.StoreKey, schema *types.Schema) Store {
	return Store{
		cdc:      cdc,
		storeKey: storeKey,

		schema:     *schema,
		stateStore: newKVStore(storeKey).Prefix([]byte{0}),
		txStore:    newKVStore(storeKey).Prefix([]byte{1}),
		cdtStore:   NewCDTStore(cdc, newKVStore(storeKey).Prefix([]byte{2}), *schema),
	}
}

func (s Store) GetInt64Store(prefix []byte) Int64Store {
	ct, found := s.schema.MatchType(prefix)
	if !found {
		panic("prefix not found")
	}
	if ct != types.CDT_TYPE_INT64 {
		panic(fmt.Sprintf("type mismatch: %v != %v", ct, types.CDT_TYPE_INT64))
	}
	return NewInt64Store(s.schema, s.stateStore, s.cdtStore, prefix)
}

func (s Store) GetGSetStore(prefix []byte) GSetStore {
	ct, found := s.schema.MatchType(prefix)
	if !found {
		panic("prefix not found")
	}
	if ct != types.CDT_TYPE_GSET {
		panic(fmt.Sprintf("type mismatch: %v != %v", ct, types.CDT_TYPE_GSET))
	}
	return NewGSetStore(s.cdc, s.schema, s.stateStore, s.cdtStore, prefix)
}

func (s Store) Precommit(ctx sdk.Context, id []byte) error {
	if s.txStore.Has(ctx, id) {
		return fmt.Errorf("id '%x' already exists", id)
	}
	ops := types.OPManagerFromContext(ctx.Context()).OPs()
	bz, err := proto.Marshal(types.PackOPs(ops))
	if err != nil {
		return err
	}
	s.txStore.Set(ctx, id, bz)
	for _, op := range ops {
		if err := s.cdtStore.ApplyOP(ctx, op); err != nil {
			return err
		}
	}
	return nil
}

func (s Store) Commit(ctx sdk.Context, id []byte) error {
	bz := s.txStore.Get(ctx, id)
	if bz == nil {
		return fmt.Errorf("id '%x' not found", id)
	}
	var anyOPs types.AnyOPs
	if err := proto.Unmarshal(bz, &anyOPs); err != nil {
		return err
	}
	ops, err := types.UnpackOPs(s.cdc, anyOPs.Ops)
	if err != nil {
		return err
	}
	if err := s.apply(ctx, ops); err != nil {
		return err
	}
	if err := s.clean(ctx, id, ops); err != nil {
		return err
	}
	return nil
}

func (s Store) Abort(ctx sdk.Context, id []byte) error {
	bz := s.txStore.Get(ctx, id)
	if bz == nil {
		// NOTE: unknown id may be indicates the aborted transaction
		return nil
	}
	var anyOPs types.AnyOPs
	if err := proto.Unmarshal(bz, &anyOPs); err != nil {
		return err
	}
	ops, err := types.UnpackOPs(s.cdc, anyOPs.Ops)
	if err != nil {
		return err
	}
	s.clean(ctx, id, ops)
	return nil
}

func (s Store) CommitImmediately(ctx sdk.Context) {
	ops := types.OPManagerFromContext(ctx.Context()).OPs()
	s.apply(ctx, ops)
}

func (s Store) apply(ctx sdk.Context, ops []types.OP) error {
	for _, op := range ops {
		// refactoring this?
		switch op := op.(type) {
		case types.Int64OP:
			v := s.stateStore.GetInt64(ctx, op.Key())
			s.stateStore.SetInt64(ctx, op.Key(), op.AddTo(v))
		case *types.GSetOP:
			set := types.NewGSetValueFromBytes(s.stateStore.Get(ctx, op.Key()))
			set.Add(op.Add...)
			s.stateStore.Set(ctx, op.Key(), set.Bytes())
		default:
			panic(fmt.Sprintf("unknown op: %T", op))
		}
	}
	return nil
}

func (s Store) clean(ctx sdk.Context, id []byte, ops []types.OP) error {
	if !s.txStore.Has(ctx, id) {
		panic(fmt.Errorf("id '%x' not found", id))
	}
	s.txStore.Delete(ctx, id)
	for _, op := range ops {
		if err := s.cdtStore.ApplyOP(ctx, op.Inverse()); err != nil {
			return err
		}
	}
	return nil
}

func buildKey(prefix []byte, key []byte) []byte {
	if prefix == nil {
		return key
	}
	k := make([]byte, len(prefix)+len(key))
	copy(k[:len(prefix)], prefix)
	copy(k[len(prefix):], key)
	return k
}
