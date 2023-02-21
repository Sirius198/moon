package ibank

import (
	"moon/x/ibank/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	height := ctx.BlockHeight()
	var _ = height
}
