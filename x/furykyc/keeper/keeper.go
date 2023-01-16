package keeper

import (
	"fmt"
	"time"

	mint "github.com/tessornetwork/fury/x/furymint/keeper"
	government "github.com/tessornetwork/fury/x/government/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/tessornetwork/fury/x/furykyc/types"
)

const (
	secondsPerYear time.Duration = time.Hour * 24 * 365
)

var membershipCosts = map[string]int64{
	types.MembershipTypeGreen:  1,
	types.MembershipTypeBronze: 25,
	types.MembershipTypeSilver: 250,
	types.MembershipTypeGold:   2500,
	types.MembershipTypeBlack:  50000,
}

type Keeper struct {
	cdc           codec.Codec
	storeKey      sdk.StoreKey
	memKey        sdk.StoreKey
	bankKeeper    bank.Keeper
	GovKeeper     government.Keeper
	accountKeeper auth.AccountKeeper
	mintKeeper    mint.Keeper
}

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey sdk.StoreKey,
	bankKeeper bank.Keeper,
	govKeeper government.Keeper,
	accountKeeper auth.AccountKeeper,
	mintKeeper mint.Keeper,
) *Keeper {
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		bankKeeper:    bankKeeper,
		GovKeeper:     govKeeper,
		accountKeeper: accountKeeper,
		mintKeeper:    mintKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
