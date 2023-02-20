package moon_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "moon/testutil/keeper"
	"moon/testutil/nullify"
	"moon/x/moon"
	"moon/x/moon/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MoonKeeper(t)
	moon.InitGenesis(ctx, *k, genesisState)
	got := moon.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
