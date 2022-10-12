package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	gammtypes "github.com/osmosis-labs/osmosis/v12/x/gamm/types"
	"github.com/osmosis-labs/osmosis/v12/x/protorev/types"
)

type Hooks struct {
	k Keeper
}

var (
	_ gammtypes.GammHooks = Hooks{}
)

// Create new pool incentives hooks.
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// AfterPoolCreated creates a gauge for each pool’s lockable duration.
func (h Hooks) AfterPoolCreated(ctx sdk.Context, sender sdk.AccAddress, poolId uint64) {
	fmt.Println("AfterPoolCreated: IN THE HOOK")

	pool, err := h.k.gammKeeper.GetPoolAndPoke(ctx, poolId)

	if err != nil {
		panic(err)
	}

	fmt.Println(pool.GetTotalPoolLiquidity(ctx))

	var allDenoms []string

	for _, coin := range pool.GetTotalPoolLiquidity(ctx) {
		allDenoms = append(allDenoms, coin.Denom)
	}

	h.k.UpdateConnectedTokens(ctx, &allDenoms)

	h.k.UpdateConnectedTokensToPoolIDs(ctx, allDenoms, poolId)

}

// AfterJoinPool hook is a noop.
func (h Hooks) AfterJoinPool(ctx sdk.Context, sender sdk.AccAddress, poolId uint64, enterCoins sdk.Coins, shareOutAmount sdk.Int) {
}

// AfterExitPool hook is a noop.
func (h Hooks) AfterExitPool(ctx sdk.Context, sender sdk.AccAddress, poolId uint64, shareInAmount sdk.Int, exitCoins sdk.Coins) {
}

// AfterSwap hook is a noop.
func (h Hooks) AfterSwap(ctx sdk.Context, sender sdk.AccAddress, poolId uint64, input sdk.Coins, output sdk.Coins) {

	needToArb := types.NeedToArb{NeedToArb: true}
	h.k.SetNeedToArb(ctx, &needToArb)

}
