package keeper

import (
	"testing"

	v300 "github.com/tessornetwork/fury/x/government/legacy/v3.0.0"
	"github.com/tessornetwork/fury/x/government/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

var governmentTestAddress, _ = sdk.AccAddressFromBech32("cosmos1lwmppctrr6ssnrmuyzu554dzf50apkfvd53jx0")
var notGovernmentAddress, _ = sdk.AccAddressFromBech32("cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae")

// This function creates an environment to test the government module
// if address is defined it will be used to add the government address
func setupKeeperWithGovernmentAddress(t testing.TB, address sdk.AccAddress) (*Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	keeper := NewKeeper(
		codec.NewProtoCodec(registry), storeKey, memStoreKey,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	if address != nil {
		store := ctx.KVStore(keeper.storeKey)
		store.Set([]byte(types.GovernmentStoreKey), address)
	}

	return keeper, ctx
}

func setupKeeperWithV300Government(t testing.TB, address sdk.AccAddress) (*Keeper, sdk.Context) {
	k, ctx := setupKeeperWithGovernmentAddress(t, nil)

	require.NotNil(t, address)

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(v300.GovernmentStoreKey), address)

	return k, ctx
}
