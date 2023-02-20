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

		TransactionList: []types.Transaction{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TransactionCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IbankKeeper(t)
	ibank.InitGenesis(ctx, *k, genesisState)
	got := ibank.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TransactionList, got.TransactionList)
	require.Equal(t, genesisState.TransactionCount, got.TransactionCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
