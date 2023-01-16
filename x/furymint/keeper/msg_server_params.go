package keeper

import (
	"context"
	"fmt"

	"github.com/tessornetwork/fury/x/furymint/types"
	ctypes "github.com/tessornetwork/fury/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetParams(goCtx context.Context, msg *types.MsgSetParams) (*types.MsgSetParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	gov := k.govKeeper.GetGovernmentAddress(ctx)
	msgGovAddr, e := sdk.AccAddressFromBech32(msg.Signer)
	if e != nil {
		return nil, e
	}
	if !(gov.Equals(msgGovAddr)) {
		return nil, sdkErr.Wrap(sdkErr.ErrUnauthorized, fmt.Sprintf("%s cannot set params", msg.Signer))
	}

	if err := msg.Params.Validate(); err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrInvalidRequest, err.Error())
	}

	k.UpdateParams(ctx, *msg.Params)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		eventSetParams,
		sdk.NewAttribute("params", msg.Params.String()),
	))
	ctypes.EmitCommonEvents(ctx, msg.Signer)

	logger := k.Logger(ctx)
	logger.Debug("Params successfully set up")

	return &types.MsgSetParamsResponse{}, nil
}
