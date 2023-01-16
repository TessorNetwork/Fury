package v400

import (
	v300 "github.com/tessornetwork/fury/x/government/legacy/v3.0.0"
	"github.com/tessornetwork/fury/x/government/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	v300key := []byte(v300.GovernmentStoreKey)

	if !store.Has(v300key) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Gov address not present")
	}
	migrateGovKeys(store)
	return nil
}

func migrateGovKeys(store sdk.KVStore) {
	v300key := []byte(v300.GovernmentStoreKey)
	govValue := store.Get(v300key)
	store.Set([]byte(types.GovernmentStoreKey), govValue)
	store.Delete(v300key)
}
