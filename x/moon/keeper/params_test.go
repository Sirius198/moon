package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "moon/testutil/keeper"
	"moon/x/moon/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MoonKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
