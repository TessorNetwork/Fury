package keeper

import (
	"context"

	"github.com/tessornetwork/fury/x/furymint/types"
	ctypes "github.com/tessornetwork/fury/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MintFUSD(goCtx context.Context, msg *types.MsgMintFUSD) (*types.MsgMintFUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var requestCoins sdk.Coins
	for _, coin := range msg.DepositAmount {
		requestCoins = append(requestCoins, *coin)
	}

	err := k.NewPosition(
		ctx,
		msg.Depositor,
		requestCoins,
		msg.ID,
	)
	if err != nil {
		return nil, errors.Wrap(errors.ErrInvalidRequest, err.Error())
	}
	ctypes.EmitCommonEvents(ctx, msg.Depositor)
	return &types.MsgMintFUSDResponse{
		ID: msg.ID,
	}, nil
}

func (k msgServer) BurnFUSD(goCtx context.Context, msg *types.MsgBurnFUSD) (*types.MsgBurnFUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}
	residualAmount, err := k.RemoveFUSD(
		ctx,
		signer,
		msg.ID,
		*msg.Amount,
	)
	if err != nil {
		return nil, err
	}
	ctypes.EmitCommonEvents(ctx, msg.Signer)
	residualCredits := sdk.NewCoin(types.CreditsDenom, residualAmount)

	return &types.MsgBurnFUSDResponse{
		ID:       msg.ID,
		Residual: &residualCredits,
	}, nil
}
