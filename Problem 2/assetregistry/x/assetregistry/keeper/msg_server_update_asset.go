package keeper

import (
	"context"
	"fmt"

	"assetregistry/x/assetregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAsset(goCtx context.Context, msg *types.MsgUpdateAsset) (*types.MsgUpdateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var asset = types.Asset{
		Id:          msg.Id,
		Name:        msg.Name,
		Description: msg.Description,
		Owner:       msg.Owner,
		Value:       msg.Value,
	}
	val, found := k.GetAsset(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("asset with ID %d doesn't exist", msg.Id))
	}
	if msg.Owner != val.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetAsset(ctx, asset)
	return &types.MsgUpdateAssetResponse{}, nil
}
