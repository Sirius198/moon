package ibank_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "moon/testutil/keeper"
	"moon/testutil/nullify"
	"moon/x/ibank"
	"moon/x/ibank/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IbankKeeper(t)
	ibank.InitGenesis(ctx, *k, genesisState)
	got := ibank.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
