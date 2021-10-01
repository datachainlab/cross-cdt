package keeper

import (
	"fmt"

	cdtkeeper "github.com/datachainlab/cross-cdt/x/cdt/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	store cdtkeeper.Int64Store
}

func NewKeeper(store cdtkeeper.Int64Store) Keeper {
	return Keeper{store: store}
}

func (k Keeper) Mint(ctx sdk.Context, account string, amount int64) error {
	assertAmount(amount)
	accKey, tsKey := accountKey(account), totalSupplyKey()
	k.store.Add(ctx, accountKey(account), amount)
	k.store.Add(ctx, totalSupplyKey(), amount)
	return k.checkBalanceInvariant(ctx, accKey, tsKey)
}

func (k Keeper) Burn(ctx sdk.Context, account string, amount int64) error {
	assertAmount(amount)
	accKey, tsKey := accountKey(account), totalSupplyKey()
	k.store.Add(ctx, accKey, -amount)
	k.store.Add(ctx, tsKey, -amount)
	return k.checkBalanceInvariant(ctx, accKey, tsKey)
}

func (k Keeper) Transfer(ctx sdk.Context, spender, recipient string, amount int64) error {
	assertAmount(amount)
	spenderKey, recipientKey := accountKey(spender), accountKey(recipient)
	k.store.Add(ctx, spenderKey, -amount)
	k.store.Add(ctx, recipientKey, amount)
	return k.checkBalanceInvariant(ctx, spenderKey, recipientKey)
}

func (k Keeper) BalanceOf(ctx sdk.Context, account string) (int64, error) {
	return k.store.Get(ctx, accountKey(account)), nil
}

func (k Keeper) TotalSupply(ctx sdk.Context) (int64, error) {
	return k.store.Get(ctx, totalSupplyKey()), nil
}

func (k Keeper) Approve(ctx sdk.Context, owner string, spender string, amount int64) (err error) {
	assertAmount(amount)
	defer func() {
		// recover from a panic such as indefinite error
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	k.store.Set(ctx, allowanceKey(owner, spender), amount)
	return nil
}

func (k Keeper) Allowance(ctx sdk.Context, owner string, spender string) (int64, error) {
	return k.store.Get(ctx, allowanceKey(owner, spender)), nil
}

func (k Keeper) TransferFrom(ctx sdk.Context, owner string, spender string, recipient string, amount int64) error {
	assertAmount(amount)
	alwKey, ownerKey, recipientKey := allowanceKey(owner, spender), accountKey(owner), accountKey(recipient)
	k.store.Add(ctx, alwKey, -amount)
	k.store.Add(ctx, ownerKey, -amount)
	k.store.Add(ctx, recipientKey, amount)
	return k.checkBalanceInvariant(ctx, alwKey, ownerKey, recipientKey)
}

func (k Keeper) checkBalanceInvariant(ctx sdk.Context, keys ...[]byte) (err error) {
	defer func() {
		// recover from a panic such as indefinite error
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	for _, key := range keys {
		if !k.store.GTE(ctx, key, 0) {
			return fmt.Errorf("violates the balance invariant: key=%x", key)
		}
	}
	return nil
}

func accountKey(account string) []byte {
	return []byte(fmt.Sprintf("acc/%v", account))
}

func assertAmount(amount int64) {
	if amount <= 0 {
		panic("amount must be greater than 0")
	}
}

func allowanceKey(owner, spender string) []byte {
	return []byte(fmt.Sprintf("alw/%v/%v", owner, spender))
}

func totalSupplyKey() []byte {
	return []byte("tsp")
}
