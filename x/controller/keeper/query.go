package keeper

import (
	"controller/x/controller/types"
)

var _ types.QueryServer = Keeper{}
