package keeper_test

import (
	"testing"

	testkeeper "alexey/testutil/keeper"
	"alexey/x/alexey/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AlexeyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
