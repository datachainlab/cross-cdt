package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/datachainlab/cross-cdt/x/cdt/testutil"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestKVStore(t *testing.T) {
	require := require.New(t)

	stk := sdk.NewKVStoreKey("state")
	s := newKVStore(stk)

	cms := testutil.MakeCMStore(t, stk)
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, tmlog.NewNopLogger())

	key0, value0 := []byte("key0"), []byte("value0")

	s.Set(ctx, key0, value0)
	require.Equal(value0, s.Get(ctx, key0))

	s1 := s.Prefix([]byte("/1/"))
	require.Nil(s1.Get(ctx, key0))

	key1, value1 := []byte("key1"), []byte("value1")

	s1.Set(ctx, key1, value1)
	require.Equal(value1, s1.Get(ctx, key1))
	require.Equal(value1, s.Get(ctx, []byte("/1/key1")))
	s1.Delete(ctx, key1)
	require.Nil(s1.Get(ctx, key1))
	require.Nil(s.Get(ctx, []byte("/1/key1")))
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
