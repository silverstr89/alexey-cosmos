package keeper_test

import (
	"context"
	"testing"

	keepertest "alexey/testutil/keeper"
	"alexey/x/alexey/keeper"
	"alexey/x/alexey/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AlexeyKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
