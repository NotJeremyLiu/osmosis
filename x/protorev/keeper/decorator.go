package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ProtoRevDecorator struct {
	ProtoRevKeeper Keeper
}

func NewProtoRevDecorator(protoRevDecorator Keeper) ProtoRevDecorator {
	return ProtoRevDecorator{
		ProtoRevKeeper: protoRevDecorator,
	}
}

func (protoRevDec ProtoRevDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {

	if !ctx.IsCheckTx() && !simulate {
		fmt.Println("Antehandler: PRINT PRINT")
	}

	fmt.Println(ctx.Context())
	fmt.Println(ctx.Logger())
	fmt.Println(ctx.EventManager().Events())
	fmt.Println(tx.GetMsgs())

	return next(ctx, tx, simulate)
}
