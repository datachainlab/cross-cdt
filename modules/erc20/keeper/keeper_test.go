package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/datachainlab/cross-cdt/modules/erc20/types"
	cdtkeeper "github.com/datachainlab/cross-cdt/x/cdt/keeper"
	"github.com/datachainlab/cross-cdt/x/cdt/testutil"
	cdttypes "github.com/datachainlab/cross-cdt/x/cdt/types"
	"github.com/stretchr/testify/require"
)

const (
	alice   = "alice"
	bob     = "bob"
	charlie = "charlie"
)

func TestERC20(t *testing.T) {
	var (
		Mint = func(account string, amount int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				store.(Keeper).Mint(ctx, account, amount)
			}
		}
		BalanceOf = func(account string, expected int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				balance, err := store.(Keeper).BalanceOf(ctx, account)
				require.NoError(t, err)
				require.Equal(t, expected, balance)
			}
		}
		Transfer = func(spender, recipient string, amount int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				require.NoError(t, store.(Keeper).Transfer(ctx, spender, recipient, amount))
			}
		}
		TransferFailed = func(spender, recipient string, amount int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				require.Error(t, store.(Keeper).Transfer(ctx, spender, recipient, amount))
			}
		}
		Approve = func(owner, spender string, amount int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				if err := store.(Keeper).Approve(ctx, owner, spender, amount); err != nil {
					panic(err)
				}
			}
		}
		Allowance = func(owner, spender string, expected int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				allowance, err := store.(Keeper).Allowance(ctx, owner, spender)
				require.NoError(t, err)
				require.Equal(t, expected, allowance)
			}
		}
		TransferFrom = func(caller, spender, recipient string, amount int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				require.NoError(t, store.(Keeper).TransferFrom(ctx, caller, spender, recipient, amount))
			}
		}
		TransferFromFailed = func(caller, spender, recipient string, amount int64) func(ctx sdk.Context, store cdttypes.StoreI) {
			return func(ctx sdk.Context, store cdttypes.StoreI) {
				require.Error(t, store.(Keeper).TransferFrom(ctx, caller, spender, recipient, amount))
			}
		}
	)

	stk := sdk.NewKVStoreKey(types.StoreKey)
	schema := cdttypes.NewSchema()
	schema.Set([]byte("/"), cdttypes.CDT_TYPE_INT64)

	var cases = []struct {
		name     string
		commands []Command
	}{
		{
			name: "single-chain: mint",
			commands: []Command{
				Commit(Mint(alice, 100)),
				Query(BalanceOf(alice, 100)),
				Commit(Mint(alice, 200)),
				Query(BalanceOf(alice, 300)),
			},
		},
		{
			name: "single-chain: transfer",
			commands: []Command{
				Commit(Mint(alice, 100)),
				Commit(Transfer(alice, bob, 100)),
				Query(BalanceOf(alice, 0), BalanceOf(bob, 100)),
			},
		},
		{
			name: "single-chain: failed to transfer",
			commands: []Command{
				Commit(Mint(alice, 100)),
				Commit(TransferFailed(alice, bob, 101)),
			},
		},
		{
			name: "single-chain: approve-transferFrom",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				Query(Allowance(alice, bob, 20)),
				Commit(TransferFrom(alice, bob, charlie, 20)),
				Query(Allowance(alice, bob, 0), BalanceOf(alice, 80), BalanceOf(bob, 0), BalanceOf(charlie, 20)),
			},
		},
		{
			name: "single-chain: approve-transferFrom",
			commands: []Command{
				Commit(Approve(alice, bob, 10)),
				Commit(Approve(alice, bob, 20)),
				Query(Allowance(alice, bob, 20)),
			},
		},
		{
			name: "single-chain: approve-transferFrom - insufficient allowance",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				Query(Allowance(alice, bob, 20)),
				Commit(TransferFromFailed(alice, bob, charlie, 21)),
			},
		},
		{
			name: "concurrent transfer - 1",
			commands: []Command{
				Commit(Mint(alice, 100)),
				AtomicPrepare(1, Transfer(alice, bob, 20)),
				AtomicPrepare(2, Transfer(alice, charlie, 30)),
				AtomicCommit(1), AtomicCommit(2),
				Query(BalanceOf(alice, 50), BalanceOf(bob, 20), BalanceOf(charlie, 30)),
			},
		},
		{
			name: "concurrent transfer - 2",
			commands: []Command{
				Commit(Mint(alice, 100)),
				AtomicPrepare(1, Transfer(alice, bob, 20)),
				AtomicPrepare(2, Transfer(alice, bob, 30)),
				AtomicCommit(1), AtomicCommit(2),
				Query(BalanceOf(alice, 50), BalanceOf(bob, 50)),
			},
		},
		{
			name: "concurrent transfer - 3",
			commands: []Command{
				Commit(Mint(alice, 10), Mint(bob, 10)),
				AtomicPrepare(1, Transfer(alice, bob, 10)),
				AtomicPrepare(2, Transfer(bob, charlie, 10)),
				AtomicCommit(1),
				Query(BalanceOf(alice, 0)),
				AtomicCommit(2),
				Query(BalanceOf(alice, 0), BalanceOf(bob, 10), BalanceOf(charlie, 10)),
			},
		},
		{
			name: "concurrent transfer - insufficient balance - 1",
			commands: []Command{
				Commit(Mint(alice, 100)),
				AtomicPrepare(1, Transfer(alice, bob, 80)),
				AtomicPrepare(2, TransferFailed(alice, charlie, 30)),
			},
		},
		{
			name: "concurrent transfer - insufficient balance - 2",
			commands: []Command{
				Commit(Mint(alice, 10), Mint(bob, 10)),
				AtomicPrepare(1, Transfer(alice, bob, 10)),
				AtomicPrepare(2, TransferFailed(bob, charlie, 20)),
			},
		},
		{
			name: "concurrent transferFrom",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				AtomicPrepare(1, TransferFrom(alice, bob, charlie, 15)),
				AtomicPrepare(2, TransferFrom(alice, bob, charlie, 5)),
				AtomicCommit(1), AtomicCommit(2),
				Query(Allowance(alice, bob, 0), BalanceOf(alice, 80), BalanceOf(bob, 0), BalanceOf(charlie, 20)),
			},
		},
		{
			name: "concurrent transferFrom - insufficient balance - 1",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				AtomicPrepare(1, TransferFromFailed(alice, bob, charlie, 21)),
			},
		},
		{
			name: "concurrent transferFrom - insufficient balance - 1",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				AtomicPrepare(1, TransferFrom(alice, bob, charlie, 15)),
				AtomicPrepare(2, TransferFromFailed(alice, bob, charlie, 6)),
			},
		},
		{
			name: "concurrent transferFrom and approve - 1",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				AtomicPrepare(1, TransferFrom(alice, bob, charlie, 15)),
				AtomicPrepare(2, ExpectErrIndefiniteState(t, Approve(alice, bob, 10))),
			},
		},
		{
			name: "concurrent transferFrom and approve - 2",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				AtomicPrepare(1, Approve(alice, bob, 15)),
				AtomicPrepare(2, TransferFrom(alice, bob, charlie, 15)),
			},
		},
		{
			name: "concurrent transferFrom and approve - 3",
			commands: []Command{
				Commit(Mint(alice, 100), Approve(alice, bob, 20)),
				AtomicPrepare(1, Approve(alice, bob, 14)),
				AtomicPrepare(2, TransferFromFailed(alice, bob, charlie, 15)),
			},
		},
		{
			name: "approve same amount",
			commands: []Command{
				Commit(Approve(alice, bob, 10)),
				AtomicPrepare(1, Approve(alice, bob, 10)),
				AtomicCommit(1),
				Query(Allowance(alice, bob, 10)),
			},
		},
		{
			name: "concurrent approve is failed",
			commands: []Command{
				AtomicPrepare(1, Approve(alice, bob, 10)),
				AtomicPrepare(2, ExpectErrIndefiniteState(t, Approve(alice, bob, 10))),
				AtomicPrepare(3, ExpectErrIndefiniteState(t, Approve(alice, bob, 20))),
			},
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			cms := testutil.MakeCMStore(t, stk)
			st := cdtkeeper.NewStore(testutil.MakeCodec(), stk, schema)
			g := testutil.NewCommandGenerater(t, st, cms)
			k := NewKeeper(st.GetInt64Store([]byte("/")))
			for _, cmd := range cs.commands {
				cmd(g)(k)
			}
		})
	}
}

// alias for testing

type (
	Command = testutil.Command
)

var (
	Commit                   = testutil.Commit
	Query                    = testutil.Query
	AtomicPrepare            = testutil.AtomicPrepare
	AtomicCommit             = testutil.AtomicCommit
	ExpectErrIndefiniteState = testutil.ExpectErrIndefiniteState
	K, V                     = testutil.K, testutil.V
)
