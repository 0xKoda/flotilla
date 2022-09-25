package keeper

import (
	"flotilla/x/dex/types"
)

var _ types.QueryServer = Keeper{}
