package keeper

import (
	"fmt"
	crosstypes "github.com/datachainlab/cross/x/core/types"

	"github.com/datachainlab/cross-cdt/x/cdt/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type kvStore struct {
	// TODO: storeKey, prefixを取得して表示してみる
	storeKey sdk.StoreKey
	prefix   []byte
}

var _ types.KVStoreI = (*kvStore)(nil)

func newKVStore(storeKey sdk.StoreKey) types.KVStoreI {
	return &kvStore{storeKey: storeKey}
}

func (s kvStore) Prefix(prefix []byte) types.KVStoreI {
	p := make([]byte, len(s.prefix)+len(prefix))
	copy(p[0:len(s.prefix)], s.prefix)
	copy(p[len(s.prefix):], prefix)
	s.prefix = p
	return s
}

func (s kvStore) KVStore(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(s.store(ctx), s.prefix)
}

func (s kvStore) Set(ctx sdk.Context, key, value []byte) {
	s.KVStore(ctx).Set(key, value)
}

func (s kvStore) SetInt64(ctx sdk.Context, key []byte, value int64) {
	s.KVStore(ctx).Set(key, sdk.Uint64ToBigEndian(uint64(value)))
}

func (s kvStore) Get(ctx sdk.Context, key []byte) []byte {
	return s.KVStore(ctx).Get(key)
}

func (s kvStore) GetInt64(ctx sdk.Context, key []byte) int64 {
	v := s.KVStore(ctx).Get(key)
	if v == nil {
		return 0
	} else if len(v) != 8 {
		panic(fmt.Errorf("got unexpected value: '%X'", v))
	}
	return int64(sdk.BigEndianToUint64(v))
}

func (s kvStore) Has(ctx sdk.Context, key []byte) bool {
	return s.KVStore(ctx).Has(key)
}

func (s kvStore) Delete(ctx sdk.Context, key []byte) {
	fmt.Printf("in Delete: %x\n", string(key))
	s.KVStore(ctx).Delete(key)
}

func (s kvStore) GetLog(ctx sdk.Context) string {
	return s.storeKey.String() + " : " + string(s.prefix)
}

func (s kvStore) store(ctx sdk.Context) sdk.KVStore {
	switch storeKey := s.storeKey.(type) {
	case *crosstypes.PrefixStoreKey:
		return prefix.NewStore(ctx.KVStore(storeKey.StoreKey), storeKey.Prefix)
	default:
		return ctx.KVStore(s.storeKey)
	}
}
