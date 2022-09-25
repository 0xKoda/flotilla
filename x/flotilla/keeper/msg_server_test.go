package keeper_test

import (
	"context"
	"testing"

	keepertest "flotilla/testutil/keeper"
	"flotilla/x/flotilla/keeper"
	"flotilla/x/flotilla/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FlotillaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
