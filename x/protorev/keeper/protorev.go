package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v12/x/protorev/types"
)

func (k Keeper) GetNeedToArb(ctx sdk.Context) *types.NeedToArb {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyNeedToArb)
	value := store.Get(types.KeyNeedToArb)

	if len(value) == 0 {
		return nil
	}

	ret := &types.NeedToArb{}
	ret.Unmarshal(value)
	return ret
}
