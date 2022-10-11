package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
}

type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type GAMMKeeper interface {
	GetNextPoolId(ctx sdk.Context) uint64
}
