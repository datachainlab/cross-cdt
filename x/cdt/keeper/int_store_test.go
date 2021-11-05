package keeper

import (
	"fmt"
	"testing"

	"github.com/datachainlab/cross-cdt/x/cdt/testutil"
	"github.com/datachainlab/cross-cdt/x/cdt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestIntStore(t *testing.T) {
	var (
		Get = func(k []byte, expectedValue int64) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				v := store.(Int64Store).Get(ctx, []byte(k))
				if expectedValue != v {
					panic(fmt.Sprintf("%v != %v", expectedValue, v))
				}
			}
		}
		Add = func(k []byte, v int64) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				store.(Int64Store).Add(ctx, []byte(k), v)
			}
		}
		GT = func(k []byte, v int64, expected bool) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				require.Equal(t, expected, store.(Int64Store).GT(ctx, []byte(k), v))
			}
		}
		GTE = func(k []byte, v int64, expected bool) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				require.Equal(t, expected, store.(Int64Store).GTE(ctx, []byte(k), v))
			}
		}
		LT = func(k []byte, v int64, expected bool) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				require.Equal(t, expected, store.(Int64Store).LT(ctx, []byte(k), v))
			}
		}
		LTE = func(k []byte, v int64, expected bool) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				require.Equal(t, expected, store.(Int64Store).LTE(ctx, []byte(k), v))
			}
		}
	)

	stk := sdk.NewKVStoreKey("main")
	schema := types.NewSchema()
	schema.Set([]byte("test/"), types.CDT_TYPE_INT64)

	var cases = []struct {
		name     string
		commands []Command
	}{
		{
			name: "single-chain: set-get",
			commands: []Command{
				Commit(Add(K(0), 1), Get(K(0), 1)),
				Query(Get(K(0), 1)),
			},
		},
		{
			name: "single-chain: add-get",
			commands: []Command{
				Commit(Add(K(0), 1), Get(K(0), 1)),
				Query(Get(K(0), 1)),
			},
		},
		{
			name: "single-chain: add-sub-get",
			commands: []Command{
				Commit(Add(K(0), 2), Add(K(0), -1), Get(K(0), 1)),
				Query(Get(K(0), 1)),
			},
		},
		{
			name: "single-chain: set-add-get",
			commands: []Command{
				Commit(Add(K(0), 1), Add(K(0), 1), Get(K(0), 2)),
				Query(Get(K(0), 2)),
			},
		},
		{
			name: "Add accesses same key in concurrency",
			commands: []Command{
				AtomicPrepare(1, Add(K(0), 1)),
				AtomicPrepare(2, Add(K(0), 2)),
				AtomicCommit(1),
				AtomicCommit(2),
				Query(Get(K(0), 3)),
			},
		},
		{
			name: "No operations",
			commands: []Command{
				AtomicPrepare(1),
				AtomicCommit(1),
			},
		},
		{
			name: "No operations but aborted",
			commands: []Command{
				AtomicPrepare(1),
				AtomicAbort(1),
			},
		},
		{
			name: "Add(0) doesn't get a lock",
			commands: []Command{
				Commit(Add(K(0), 0)),
				AtomicPrepare(1, Add(K(0), 1), Add(K(0), -1)),
				Query(Get(K(0), 0)),
				AtomicCommit(1),
				Query(Get(K(0), 0)),
			},
		},
		{
			name: "Boundary value checks for compare functions",
			commands: []Command{
				Commit(Add(K(0), 1)),
				AtomicPrepare(1, Add(K(0), 1)),
				AtomicPrepare(
					2,
					ExpectErrIndefiniteState(t, GT(K(0), 1, false)),
					GTE(K(0), 1, true),
					LT(K(0), 1, false),
					ExpectErrIndefiniteState(t, LTE(K(0), 1, false)),
				),
				AtomicCommit(1),
				Query(
					Get(K(0), 2),
					GT(K(0), 1, true),
					GTE(K(0), 1, true),
					LT(K(0), 1, false),
					LTE(K(0), 1, false),
					Get(K(0), 2),
				),
			},
		},
		{
			name: "check if a lock is released after aborted",
			commands: []Command{
				AtomicPrepare(1, Add(K(0), 1)),
				AtomicAbort(1),
				Query(Get(K(0), 0)),
				Commit(Add(K(0), 1)),
				Query(Get(K(0), 1)),
			},
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			cms := testutil.MakeCMStore(t, stk)
			st := NewStore(testutil.MakeCodec(), stk, schema)
			ist := st.GetInt64Store([]byte("test/"))
			var g = testutil.NewCommandGenerater(t, st, cms)

			for _, cmd := range cs.commands {
				cmd(g)(ist)
			}
		})
	}
}
