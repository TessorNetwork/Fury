package keeper

import (
	"github.com/tessornetwork/fury/x/furymint/types"
)

var _ types.QueryServer = Keeper{}
