package keeper

import (
	"testing"

	"github.com/datachainlab/cross-cdt/x/cdt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGSetStore(t *testing.T) {
	var (
		Get = func(k []byte, expected ...types.Element) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				require.EqualValues(t, expected, store.(GSetStore).Get(ctx, []byte(k)))
			}
		}
		Add = func(k []byte, elements ...types.Element) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				store.(GSetStore).Add(ctx, []byte(k), elements...)
			}
		}
		Lookup = func(k []byte, expected bool, elements ...types.Element) func(ctx sdk.Context, store types.StoreI) {
			return func(ctx sdk.Context, store types.StoreI) {
				require.Equal(t, expected, store.(GSetStore).Lookup(ctx, []byte(k), elements...))
			}
		}
	)

	stk := sdk.NewKVStoreKey("main")
	schema := types.NewSchema()
	schema.Set([]byte("test/"), types.CDT_TYPE_GSET)

	var cases = []struct {
		name     string
		commands []Command
	}{
		{
			name: "single-chain - 1",
			commands: []Command{
				Commit(Add(K(0), V(0))),
				Query(
					Lookup(K(0), true, V(0)),
					Get(K(0), V(0)),
				),
			},
		},
		{
			name: "single-chain - 2",
			commands: []Command{
				Commit(Add(K(0), V(0), V(1))),
				Query(
					Lookup(K(0), true, V(0)),
					Lookup(K(0), true, V(1)),
					Lookup(K(0), false, V(2)),
					Lookup(K(0), true, V(0), V(1)),
					Lookup(K(0), false, V(0), V(1), V(2)),
					Get(K(0), V(0), V(1)),
				),
			},
		},
		{
			name: "single-chain - 3",
			commands: []Command{
				Commit(Add(K(0), V(0), V(1))),
				Query(
					Lookup(K(0), true, V(0)),
					Lookup(K(0), true, V(1)),
					Lookup(K(0), true, V(0), V(1)),
					Get(K(0), V(0), V(1)),
				),
			},
		},
		{
			name: "Set does not keep duplicate elements - 1",
			commands: []Command{
				Commit(Add(K(0), V(0))),
				Commit(Add(K(0), V(1), V(2))),
				Query(
					Get(K(0), V(0), V(1), V(2)),
				),
			},
		},
		{
			name: "Set does not keep duplicate elements - 2",
			commands: []Command{
				Commit(Add(K(0), V(0))),
				AtomicPrepare(1, Add(K(0), V(0))),
				AtomicPrepare(2, Add(K(0), V(1), V(2))),
				AtomicCommit(1), AtomicCommit(2),
				Query(
					Get(K(0), V(0), V(1), V(2)),
				),
			},
		},
		{
			name: "conflict: failed to lookup the provisional state -1",
			commands: []Command{
				Commit(Add(K(0), V(0))),
				AtomicPrepare(1, Add(K(0), V(1))),
				AtomicPrepare(2, Lookup(K(0), true, V(0))),
				AtomicPrepare(3, ExpectErrIndefiniteState(t, Lookup(K(0), true, V(1)))),
			},
		},
		{
			name: "conflict: failed to lookup the provisional state - 2",
			commands: []Command{
				AtomicPrepare(1, Add(K(0), V(0))),
				AtomicPrepare(2, Add(K(0), V(1))),
				AtomicCommit(2),
				Query(
					Lookup(K(0), true, V(1)),
					ExpectErrIndefiniteState(t, Lookup(K(0), true, V(0))),
				),
				AtomicCommit(1),
				Query(
					Lookup(K(0), true, V(0), V(1)),
					Get(K(0), V(0), V(1)),
				),
			},
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			cms := makeCMStore(t, stk)
			st := NewStore(makeCodec(), stk, schema)
			gst := st.GetGSetStore([]byte("test/"))
			var g = NewCommandGenerater(t, st, cms)

			for _, cmd := range cs.commands {
				cmd(g)(gst)
			}
		})
	}
}
