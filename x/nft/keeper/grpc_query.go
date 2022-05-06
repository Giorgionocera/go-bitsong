package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/go-bitsong/x/nft/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) NFTInfo(c context.Context, req *types.QueryNFTInfoRequest) (*types.QueryNFTInfoResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	nft, err := k.GetNFTById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	metadata, err := k.GetMetadataById(ctx, nft.MetadataId)
	if err != nil {
		return nil, err
	}
	return &types.QueryNFTInfoResponse{
		Nft:      nft,
		Metadata: metadata,
	}, nil
}

func (k Keeper) Metadata(c context.Context, req *types.QueryMetadataRequest) (*types.QueryMetadataResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	metadata, err := k.GetMetadataById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.QueryMetadataResponse{
		Metadata: metadata,
	}, nil
}

func (k Keeper) Collection(c context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	collection, err := k.GetCollectionById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	nftIds := k.GetCollectionNftRecords(ctx, req.Id)
	return &types.QueryCollectionResponse{
		Collection: collection,
		NftIds:     nftIds,
	}, nil
}