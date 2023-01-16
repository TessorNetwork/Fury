package keeper

import (
	"reflect"
	"testing"

	"github.com/tessornetwork/fury/x/furymint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_msgServer_MintFUSD(t *testing.T) {

	type args struct {
		msg *types.MsgMintFUSD
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgMintFUSDResponse
		wantErr bool
	}{
		{
			name: "invalid depositor",
			args: args{
				msg: &types.MsgMintFUSD{
					Depositor:     "",
					DepositAmount: []*sdk.Coin{&validDepositCoin},
					ID:            testEtp.ID,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid coins",
			args: args{
				msg: &types.MsgMintFUSD{
					Depositor:     testEtpOwner.String(),
					DepositAmount: []*sdk.Coin{&validDepositCoin, &inValidDepositCoin},
					ID:            testEtp.ID,
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				msg: &types.MsgMintFUSD{
					Depositor:     testEtpOwner.String(),
					DepositAmount: []*sdk.Coin{&validDepositCoin},
					ID:            testID,
				},
			},
			want: &types.MsgMintFUSDResponse{ID: testEtp.ID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wctx, bk, _, _, msgServer := SetupMsgServer()

			if !tt.wantErr {
				ownerAddr, err := sdk.AccAddressFromBech32(tt.args.msg.Depositor)
				require.NoError(t, err)
				coins := sdk.NewCoins(sdk.NewInt64Coin(types.BondDenom, 200))
				ctx := sdk.UnwrapSDKContext(wctx)
				err = bk.MintCoins(ctx, types.ModuleName, coins)
				require.NoError(t, err)
				err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddr, coins)
				require.NoError(t, err)
				//err = bk.AddCoins(sdk.UnwrapSDKContext(wctx), ownerAddr, sdk.NewCoins(sdk.NewInt64Coin(types.BondDenom, 200)))
				//require.NoError(t, err)
			}

			got, err := msgServer.MintFUSD(wctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("msgServer.MintFUSD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("msgServer.MintFUSD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_BurnFUSD(t *testing.T) {

	type args struct {
		msg *types.MsgBurnFUSD
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgBurnFUSDResponse
		wantErr bool
	}{
		{
			name: "empty signer",
			args: args{
				msg: &types.MsgBurnFUSD{
					Signer: "",
					Amount: &validBurnCoin,
					ID:     testID,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid coins",
			args: args{
				msg: &types.MsgBurnFUSD{
					Signer: testEtpOwner.String(),
					Amount: &inValidBurnCoin,
					ID:     testID,
				},
			},
			wantErr: true,
		},
		// TODO: fix liquidity pool setup
		/*{
			name: "ok",
			args: args{
				msg: &types.MsgBurnFUSD{
					Signer: testEtpOwner.String(),
					Amount: testEtp.Credits,
					ID:     testID,
				},
			},
			want: &types.MsgBurnFUSDResponse{ID: testEtp.ID, Residual: &zeroUFUSD}, // TODO check residuals
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wctx, bk, _, k, msgServer := SetupMsgServer()

			if !tt.wantErr {
				ctx := sdk.UnwrapSDKContext(wctx)
				k.SetPosition(ctx, testEtp)
				_ = bk.MintCoins(ctx, types.ModuleName, testLiquidityPool)
				_ = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, testEtpOwner, sdk.NewCoins(*testEtp.Credits))
				//_ = bk.AddCoins(sdk.UnwrapSDKContext(wctx), testEtpOwner, sdk.NewCoins(*testEtp.Credits))
			}

			got, err := msgServer.BurnFUSD(wctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("msgServer.BurnFUSD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("msgServer.BurnFUSD() = %v, want %v", got, tt.want)
			}
		})
	}
}
