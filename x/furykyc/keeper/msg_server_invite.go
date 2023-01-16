package keeper

import (
	"context"

	"github.com/tessornetwork/fury/x/furykyc/types"
	ctypes "github.com/tessornetwork/fury/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"
)

// InviteUser handle message MsgInviteUser
func (k msgServer) InviteUser(goCtx context.Context, msg *types.MsgInviteUser) (*types.MsgInviteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verify that the user that is invited is not present on the chain
	msgRecipient, _ := sdk.AccAddressFromBech32(msg.Recipient)
	if k.accountKeeper.GetAccount(ctx, msgRecipient) != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnauthorized, "cannot invite existing user")
	}

	msgSender, _ := sdk.AccAddressFromBech32(msg.Sender)

	// Try inviting the user
	if err := k.SetInvite(ctx, msgRecipient, msgSender); err != nil {
		return nil, err
	}
	ctypes.EmitCommonEvents(ctx, msg.Sender)
	return &types.MsgInviteUserResponse{
		Status: "1",
	}, nil

}
