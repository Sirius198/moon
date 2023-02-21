package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"moon/x/ibank/types"
)

func (k Keeper) ShowIncoming(goCtx context.Context, req *types.QueryShowIncomingRequest) (*types.QueryShowIncomingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryShowIncomingResponse{}, nil
}
