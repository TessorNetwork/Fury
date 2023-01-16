package keeper

import (
	"testing"

	"github.com/tessornetwork/fury/x/government/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_GovernmentAddr(t *testing.T) {

	tests := []struct {
		name       string
		request    *types.QueryGovernmentAddrRequest
		government sdk.AccAddress
		wantErr    bool
	}{
		{
			"valid request",
			&types.QueryGovernmentAddrRequest{},
			governmentTestAddress,
			false,
		},
		{
			"empty request",
			nil,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, ctx := setupKeeperWithGovernmentAddress(t, tt.government)

			require.Equal(t, tt.government, k.GetGovernmentAddress(ctx))

			c := sdk.WrapSDKContext(ctx)

			got, err := k.GovernmentAddr(c, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keeper.GovernmentAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				require.Equal(t, &types.QueryGovernmentAddrResponse{GovernmentAddress: k.GetGovernmentAddress(ctx).String()}, got)
			}
		})
	}
}
