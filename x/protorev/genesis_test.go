package protorev_test

import (
	"testing"

	keepertest "github.com/osmosis-labs/osmosis/v12/testutil/keeper"
	"github.com/osmosis-labs/osmosis/v12/testutil/nullify"
	"github.com/osmosis-labs/osmosis/v12/x/protorev"
	"github.com/osmosis-labs/osmosis/v12/x/protorev/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProtorevKeeper(t)
	protorev.InitGenesis(ctx, *k, genesisState)
	got := protorev.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
