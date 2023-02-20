package keeper

import (
	"moon/x/moon/types"
)

var _ types.QueryServer = Keeper{}
