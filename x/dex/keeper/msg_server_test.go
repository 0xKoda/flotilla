package keeper_test

import (
	"context"
	"testing"

	keepertest "flotilla/testutil/keeper"
	"flotilla/x/dex/keeper"
	"flotilla/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DexKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
