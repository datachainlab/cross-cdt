package testutil

import (
	"fmt"
	"testing"

	"github.com/datachainlab/cross-cdt/x/cdt/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdkstore "github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	contracttypes "github.com/datachainlab/cross/x/core/contract/types"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"
)

/* utility functions for testing */

type StoreCommand func(store types.StoreI)

type OPCommand func(ctx sdk.Context, store types.StoreI)

type Command func(g *CommandGenerator) StoreCommand

type CommandGenerator struct {
	t   *testing.T
	st  contracttypes.CommitStoreI
	cms sdk.CommitMultiStore
}

func NewCommandGenerater(t *testing.T, st contracttypes.CommitStoreI, cms sdk.CommitMultiStore) *CommandGenerator {
	return &CommandGenerator{t: t, st: st, cms: cms}
}

func (g *CommandGenerator) AtomicPrepare(id uint64, ops ...OPCommand) StoreCommand {
	return func(store types.StoreI) {
		ctx := MakeAtomicModeContext(g.cms, types.NewOPManager())
		for _, op := range ops {
			op(ctx, store)
		}
		require.NoError(g.t, g.st.Precommit(ctx, sdk.Uint64ToBigEndian(id)))
	}
}

func (g *CommandGenerator) AtomicCommit(id uint64) StoreCommand {
	return func(store types.StoreI) {
		ctx := MakeAtomicModeContext(g.cms, types.NewOPManager())
		g.st.Commit(ctx, sdk.Uint64ToBigEndian(id))
	}
}

func (g *CommandGenerator) Commit(ops ...OPCommand) StoreCommand {
	return func(store types.StoreI) {
		ctx := MakeBasicModeContext(g.cms, types.NewOPManager())
		for _, op := range ops {
			op(ctx, store)
		}
		g.st.CommitImmediately(ctx)
	}
}

func (g *CommandGenerator) Query(ops ...OPCommand) StoreCommand {
	return func(store types.StoreI) {
		ctx, _ := MakeBasicModeContext(g.cms, types.NewOPManager()).CacheContext()
		for _, op := range ops {
			op(ctx, store)
		}
		g.st.CommitImmediately(ctx)
	}
}

func AtomicPrepare(id uint64, ops ...OPCommand) func(g *CommandGenerator) StoreCommand {
	return func(g *CommandGenerator) StoreCommand {
		return g.AtomicPrepare(id, ops...)
	}
}

func AtomicCommit(id uint64) func(g *CommandGenerator) StoreCommand {
	return func(g *CommandGenerator) StoreCommand {
		return g.AtomicCommit(id)
	}
}

func Commit(ops ...OPCommand) func(g *CommandGenerator) StoreCommand {
	return func(g *CommandGenerator) StoreCommand {
		return g.Commit(ops...)
	}
}

func Query(ops ...OPCommand) func(g *CommandGenerator) StoreCommand {
	return func(g *CommandGenerator) StoreCommand {
		return g.Query(ops...)
	}
}

func ExpectErrIndefiniteState(t *testing.T, op OPCommand) OPCommand {
	return func(ctx sdk.Context, store types.StoreI) {
		require.PanicsWithError(t, types.ErrIndefiniteState.Error(), func() {
			op(ctx, store)
		})
	}
}

func MakeContext(cms sdk.CommitMultiStore) sdk.Context {
	return sdk.NewContext(cms, tmproto.Header{}, false, tmlog.NewNopLogger())
}

func MakeBasicModeContext(cms sdk.CommitMultiStore, lkmgr types.OPManager) sdk.Context {
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, tmlog.NewNopLogger())
	ctx = ctx.WithContext(types.ContextWithOPManager(ctx.Context(), lkmgr))
	return ctx.WithContext(
		contracttypes.ContextWithContractRuntimeInfo(
			ctx.Context(),
			contracttypes.ContractRuntimeInfo{CommitMode: contracttypes.BasicMode},
		),
	)
}

func MakeAtomicModeContext(cms sdk.CommitMultiStore, lkmgr types.OPManager) sdk.Context {
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, tmlog.NewNopLogger())
	ctx = ctx.WithContext(types.ContextWithOPManager(ctx.Context(), lkmgr))
	return ctx.WithContext(
		contracttypes.ContextWithContractRuntimeInfo(
			ctx.Context(),
			contracttypes.ContractRuntimeInfo{CommitMode: contracttypes.AtomicMode},
		),
	)
}

func MakeCMStore(t *testing.T, key sdk.StoreKey) sdk.CommitMultiStore {
	require := require.New(t)
	d, err := db.NewDB("test", db.MemDBBackend, "")
	if err != nil {
		panic(err)
	}
	cms := sdkstore.NewCommitMultiStore(d)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, d)
	require.NoError(cms.LoadLatestVersion())
	return cms
}

func MakeCodec() codec.Codec {
	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	types.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}

func K(idx int) []byte {
	return []byte(fmt.Sprintf("k%v", idx))
}

func V(idx int) []byte {
	return []byte(fmt.Sprintf("v%v", idx))
}
