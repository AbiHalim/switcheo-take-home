package keeper_test

import (
  "context"
  "testing"

  "github.com/stretchr/testify/require"

  keepertest "assetregistry/testutil/keeper"
  "assetregistry/x/assetregistry/keeper"
  "assetregistry/x/assetregistry/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
  k, ctx := keepertest.AssetregistryKeeper(t)
  return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
  k, ms, ctx := setupMsgServer(t)
  require.NotNil(t, ms)
  require.NotNil(t, ctx)
  require.NotEmpty(t, k)
}

func TestCreateAsset_MinValue(t *testing.T) {
  msg := types.MsgCreateAsset{
    Creator:     "cosmos1...",
    Name:        "Low Value Asset",
    Description: "This asset has a low value",
    Value:       50, // Below the minimum value
  }

  _, err := msgServer.CreateAsset(context.Background(), &msg)
  require.Error(t, err)
  require.Contains(t, err.Error(), "value must be at least 100")
}

func TestUpdateAsset_MinValue(t *testing.T) {
  existingAsset := types.Asset{
    Creator:     "cosmos1...",
    Name:        "Existing Asset",
    Description: "An existing asset",
    Value:       300,
  }
  k.SetAsset(ctx, existingAsset)

  msg := types.MsgUpdateAsset{
    Creator:     "cosmos1...",
    Id:          existingAsset.Id,
    Name:        "Updated Asset",
    Description: "Updated description",
    Value:       50, // Below the minimum value
  }

  _, err := msgServer.UpdateAsset(context.Background(), &msg)
  require.Error(t, err)
  require.Contains(t, err.Error(), "value must be at least 100")
}
