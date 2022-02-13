package keeper

import (
	"github.com/nzee/scavenge/x/scavenge/types"
)

var _ types.QueryServer = Keeper{}
