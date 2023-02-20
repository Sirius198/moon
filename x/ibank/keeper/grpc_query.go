package keeper

import (
	"moon/x/ibank/types"
)

var _ types.QueryServer = Keeper{}
