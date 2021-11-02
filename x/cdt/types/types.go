package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// KVStoreI defines the expected key-value store
type KVStoreI interface {
	Prefix(prefix []byte) KVStoreI
	KVStore(ctx sdk.Context) sdk.KVStore

	Set(ctx sdk.Context, key, value []byte)
	Get(ctx sdk.Context, key []byte) []byte
	Has(ctx sdk.Context, key []byte) bool
	Delete(ctx sdk.Context, key []byte)

	GetInt64(ctx sdk.Context, key []byte) int64
	SetInt64(ctx sdk.Context, key []byte, value int64)

	GetLog(ctx sdk.Context) string
}

type StoreI interface{}
