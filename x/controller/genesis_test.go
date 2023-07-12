package controller_test

import (
	"testing"

	keepertest "controller/testutil/keeper"
	"controller/testutil/nullify"
	"controller/x/controller"
	"controller/x/controller/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ControllerKeeper(t)
	controller.InitGenesis(ctx, *k, genesisState)
	got := controller.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
