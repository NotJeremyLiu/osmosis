package keeper

import (
	"fmt"
	"sort"

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

func (k Keeper) SetNeedToArb(ctx sdk.Context, needToArb *types.NeedToArb) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyNeedToArb)

	value, err := needToArb.Marshal()

	if err != nil {
		panic(err)
	}

	store.Set(types.KeyNeedToArb, value)
}

func (k Keeper) GetArbDetails(ctx sdk.Context) *types.ArbDetails {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyArbDetails)
	value := store.Get(types.KeyArbDetails)

	if len(value) == 0 {
		return nil
	}

	ret := &types.ArbDetails{}
	ret.Unmarshal(value)
	return ret
}

func (k Keeper) SetArbDetails(ctx sdk.Context, arbDetails *types.ArbDetails) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyArbDetails)

	value, err := arbDetails.Marshal()

	if err != nil {
		panic(err)
	}

	store.Set(types.KeyArbDetails, value)
}

func (k Keeper) GetConnectedTokens(ctx sdk.Context, token *string) *types.ConnectedTokens {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyConnectedTokens)

	key := types.GetConnectedTokensStoreKey(token)

	value := store.Get(key)

	if len(value) == 0 {
		return nil
	}

	ret := &types.ConnectedTokens{}
	ret.Unmarshal(value)
	return ret
}

func (k Keeper) SetConnectedTokens(ctx sdk.Context, token *string, connectedTokens *types.ConnectedTokens) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyConnectedTokens)

	key := types.GetConnectedTokensStoreKey(token)

	value, err := connectedTokens.Marshal()

	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}

func (k Keeper) UpdateConnectedTokens(ctx sdk.Context, allDenoms *[]string) {

	for _, token := range *allDenoms {
		fmt.Println(token)
		connectedTokensRes := k.GetConnectedTokens(ctx, &token)
		fmt.Println(connectedTokensRes)

		var connectedTokens []string

		if connectedTokensRes == nil {
			connectedTokens = []string{}
		} else {
			connectedTokens = connectedTokensRes.Tokens
		}

		fmt.Println(connectedTokens)

		fmt.Println(allDenoms)

		for _, denom := range *allDenoms {
			if !types.Contains(connectedTokens, denom) && denom != token {
				connectedTokens = append(connectedTokens, denom)
			} else {
				fmt.Println("Already connected and/or same denom, not adding")
			}
		}

		fmt.Println(connectedTokens)

		k.SetConnectedTokens(ctx, &token, &types.ConnectedTokens{Tokens: connectedTokens})
		updatedConnectedTokenRes := k.GetConnectedTokens(ctx, &token)
		fmt.Println(&updatedConnectedTokenRes.Tokens)
	}
}

func (k Keeper) GetConnectedTokensToPoolIDs(ctx sdk.Context, tokenA string, tokenB string) *types.PairsToPoolIDs {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyConnectedTokensToPoolIDs)

	key := types.GetConnectedTokensToPoolIDsStoreKey(tokenA, tokenB)

	value := store.Get(key)

	if len(value) == 0 {
		return nil
	}

	ret := &types.PairsToPoolIDs{}
	ret.Unmarshal(value)
	return ret
}

func (k Keeper) SetConnectedTokensToPoolIDs(ctx sdk.Context, tokenA string, tokenB string, pairsToPoolIDs *types.PairsToPoolIDs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyConnectedTokensToPoolIDs)

	key := types.GetConnectedTokensToPoolIDsStoreKey(tokenA, tokenB)

	value, err := pairsToPoolIDs.Marshal()

	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}

func (k Keeper) UpdateConnectedTokensToPoolIDs(ctx sdk.Context, allDenoms []string, poolID uint64) {

	sort.Strings(allDenoms)

	fmt.Println(allDenoms)

	for i, tokenA := range allDenoms {
		for _, tokenB := range allDenoms[i+1:] {
			fmt.Println(tokenA)
			fmt.Println(tokenB)

			pairsToPoolIDsRes := k.GetConnectedTokensToPoolIDs(ctx, tokenA, tokenB)

			fmt.Println(pairsToPoolIDsRes)

			var pairsToPoolIDs []uint64

			if pairsToPoolIDsRes == nil {
				pairsToPoolIDs = []uint64{}
			} else {
				pairsToPoolIDs = pairsToPoolIDsRes.PoolIds
			}

			fmt.Println(pairsToPoolIDs)

			pairsToPoolIDs = append(pairsToPoolIDs, poolID)

			fmt.Println(pairsToPoolIDs)

			k.SetConnectedTokensToPoolIDs(ctx, tokenA, tokenB, &types.PairsToPoolIDs{PoolIds: pairsToPoolIDs})
			updatedPairsToPoolIDsRes := k.GetConnectedTokensToPoolIDs(ctx, tokenA, tokenB)
			fmt.Println(&updatedPairsToPoolIDsRes.PoolIds)
		}
	}
}

func (k Keeper) GetPoolRoutes(ctx sdk.Context, poolID uint64) *types.ListOfCyclicRoutes {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolToRoutes)

	key := types.GetPoolToRoutesStoreKey(poolID)

	value := store.Get(key)

	if len(value) == 0 {
		return nil
	}

	ret := &types.ListOfCyclicRoutes{}
	ret.Unmarshal(value)
	return ret
}

func (k Keeper) SetPoolRoutes(ctx sdk.Context, poolID uint64, listOfCyclicRoutes *types.ListOfCyclicRoutes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolToRoutes)

	key := types.GetPoolToRoutesStoreKey(poolID)

	value, err := listOfCyclicRoutes.Marshal()

	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}
