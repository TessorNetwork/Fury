package keeper

import (
	"context"

	"github.com/tessornetwork/fury/x/government/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GovernmentAddr(c context.Context, req *types.QueryGovernmentAddrRequest) (*types.QueryGovernmentAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryGovernmentAddrResponse{GovernmentAddress: k.GetGovernmentAddress(ctx).String()}, nil
}
