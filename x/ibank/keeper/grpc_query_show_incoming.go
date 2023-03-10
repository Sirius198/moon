package keeper

import (
	"context"

	"moon/x/ibank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
