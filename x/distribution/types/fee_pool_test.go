package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Canto-Network/Canto/v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestValidateGenesis(t *testing.T) {
	fp := types.InitialFeePool()
	require.Nil(t, fp.ValidateGenesis())

	fp2 := types.FeePool{CommunityPool: sdk.DecCoins{{Denom: "stake", Amount: sdk.NewDec(-1)}}}
	require.NotNil(t, fp2.ValidateGenesis())
}
