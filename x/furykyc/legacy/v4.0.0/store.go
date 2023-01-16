package v400

import (
	v300 "github.com/tessornetwork/fury/x/furykyc/legacy/v3.0.0"
	"github.com/tessornetwork/fury/x/furykyc/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	migrateMembershipKeys(store, cdc)
	return nil
}

func migrateMembershipKeys(store sdk.KVStore, cdc codec.BinaryCodec) {
	oldStore := prefix.NewStore(store, []byte(v300.MembershipsStorageKey))
	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()
	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		var membership types.Membership
		cdc.MustUnmarshal(oldStoreIter.Value(), &membership)
		addr, _ := sdk.AccAddressFromBech32(membership.Owner)
		memberKey := append([]byte(types.MembershipsStorageKey), addr.Bytes()...)

		// Set new key on store. Values don't change.
		store.Set(memberKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}
