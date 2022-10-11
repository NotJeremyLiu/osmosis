package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/osmosis-labs/osmosis/v12/testutil/keeper"
	"github.com/osmosis-labs/osmosis/v12/x/protorev/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ProtorevKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
