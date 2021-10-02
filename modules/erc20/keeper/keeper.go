package keeper

import (
	"errors"
	"fmt"

	cdtkeeper "github.com/datachainlab/cross-cdt/x/cdt/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper implements ERC20 token contract using CDT store
// Each function can be performed in concurrency under specific conditions using the operations provided by CDT.
type Keeper struct {
	store cdtkeeper.Int64Store
}

// NewKeeper creates a new keeper instance
func NewKeeper(store cdtkeeper.Int64Store) Keeper {
	return Keeper{store: store}
}

// Mint creates `amount` tokens and assigns them to `account`, increasing the total supply.
func (k Keeper) Mint(ctx sdk.Context, account string, amount int64) error {
	if err := assertAmount(amount); err != nil {
		return err
	}
	k.store.Add(ctx, accountKey(account), amount)
	k.store.Add(ctx, totalSupplyKey(), amount)
	return nil
}

// Burn Destroys `amount` tokens from `account`, reducing the total supply.
func (k Keeper) Burn(ctx sdk.Context, account string, amount int64) error {
	accKey, tsKey := accountKey(account), totalSupplyKey()
	if err := assertAmount(amount); err != nil {
		return err
	} else if err := k.assertSufficientBalance(ctx, accKey, amount); err != nil {
		return err
	} else if err := k.assertSufficientBalance(ctx, tsKey, amount); err != nil {
		return err
	}
	k.store.Add(ctx, accKey, -amount)
	k.store.Add(ctx, tsKey, -amount)
	return nil
}

// Transfer moves `amount` of tokens from `sender` to `recipient`.
func (k Keeper) Transfer(ctx sdk.Context, spender, recipient string, amount int64) error {
	spenderKey, recipientKey := accountKey(spender), accountKey(recipient)
	if err := assertAmount(amount); err != nil {
		return err
	} else if err := k.assertSufficientBalance(ctx, spenderKey, amount); err != nil {
		return err
	}
	k.store.Add(ctx, spenderKey, -amount)
	k.store.Add(ctx, recipientKey, amount)
	return nil
}

// BalanceOf returns the amount of tokens owned by `account`.
func (k Keeper) BalanceOf(ctx sdk.Context, account string) (int64, error) {
	return k.store.Get(ctx, accountKey(account)), nil
}

// TotalSupply returns the amount of tokens in existence.
func (k Keeper) TotalSupply(ctx sdk.Context) (int64, error) {
	return k.store.Get(ctx, totalSupplyKey()), nil
}

// Approve sets `amount` as the allowance of `spender` over the caller's tokens.
func (k Keeper) Approve(ctx sdk.Context, owner string, spender string, amount int64) (err error) {
	if err := assertAmount(amount); err != nil {
		return err
	}
	return k.setBalance(ctx, allowanceKey(owner, spender), amount)
}

// Allowance Returns the remaining number of tokens that `spender` will be
// allowed to spend on behalf of `owner` through {transferFrom}. This is
// zero by default.
func (k Keeper) Allowance(ctx sdk.Context, owner string, spender string) (int64, error) {
	return k.store.Get(ctx, allowanceKey(owner, spender)), nil
}

// TransferFrom moves `amount` tokens from `sender` to `recipient` using the
// allowance mechanism. `amount` is then deducted from the caller's allowance.
func (k Keeper) TransferFrom(ctx sdk.Context, owner string, spender string, recipient string, amount int64) error {
	alwKey, ownerKey, recipientKey := allowanceKey(owner, spender), accountKey(owner), accountKey(recipient)
	if err := assertAmount(amount); err != nil {
		return err
	} else if err := k.assertSufficientBalance(ctx, alwKey, amount); err != nil {
		return err
	} else if err := k.assertSufficientBalance(ctx, ownerKey, amount); err != nil {
		return err
	}
	k.store.Add(ctx, alwKey, -amount)
	k.store.Add(ctx, ownerKey, -amount)
	k.store.Add(ctx, recipientKey, amount)
	return nil
}

func (k Keeper) setBalance(ctx sdk.Context, key []byte, value int64) (err error) {
	defer func() {
		// recover from a panic such as indefinite error
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	k.store.Set(ctx, key, value)
	return nil
}

func (k Keeper) assertSufficientBalance(ctx sdk.Context, key []byte, value int64) (err error) {
	defer func() {
		// recover from a panic such as indefinite error
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	if !k.store.GTE(ctx, key, value) {
		return fmt.Errorf("violates the balance invariant: key=%x", key)
	}
	return nil
}

func accountKey(account string) []byte {
	return []byte(fmt.Sprintf("acc/%v", account))
}

func assertAmount(amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	return nil
}

func allowanceKey(owner, spender string) []byte {
	return []byte(fmt.Sprintf("alw/%v/%v", owner, spender))
}

func totalSupplyKey() []byte {
	return []byte("tsp")
}
