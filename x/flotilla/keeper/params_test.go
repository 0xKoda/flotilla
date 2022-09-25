package keeper_test

import (
	"testing"

	testkeeper "flotilla/testutil/keeper"
	"flotilla/x/flotilla/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FlotillaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
