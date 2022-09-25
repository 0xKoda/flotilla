package flotilla_test

import (
	"testing"

	keepertest "flotilla/testutil/keeper"
	"flotilla/testutil/nullify"
	"flotilla/x/flotilla"
	"flotilla/x/flotilla/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FlotillaKeeper(t)
	flotilla.InitGenesis(ctx, *k, genesisState)
	got := flotilla.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
