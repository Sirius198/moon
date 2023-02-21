package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"moon/x/ibank/types"
)

func (k Keeper) ShowOutgoing(goCtx context.Context, req *types.QueryShowOutgoingRequest) (*types.QueryShowOutgoingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryShowOutgoingResponse{}, nil
}
