package government

import (
	"github.com/tessornetwork/fury/x/government/keeper"
	"github.com/tessornetwork/fury/x/government/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	govAddr, err := sdk.AccAddressFromBech32(genState.GovernmentAddress)
	if err != nil {
		panic(err)
	}

	errSetGov := k.SetGovernmentAddress(ctx, govAddr)
	if errSetGov != nil {
		panic(errSetGov)
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.GovernmentAddress = k.GetGovernmentAddress(ctx).String()
	return genesis
}
