package epochs

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tessornetwork/fury/x/epochs/keeper"
	"github.com/tessornetwork/fury/x/epochs/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// set epoch info from genesis
	for _, epoch := range genState.Epochs {
		// when epoch counting start time is not set, set it
		if epoch.StartTime.Equal(time.Time{}) {
			epoch.StartTime = ctx.BlockTime()
		}

		// when epoch counting started and current_epoch_start_time is not set, set it
		if epoch.EpochCountingStarted && epoch.CurrentEpochStartTime.Equal(time.Time{}) {
			epoch.CurrentEpochStartTime = ctx.BlockTime()
		}

		k.SetEpochInfo(ctx, epoch)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Epochs = k.AllEpochInfos(ctx)
	return genesis
}
