package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mastervectormaster/lottery/x/lottery/keeper"
	"github.com/mastervectormaster/lottery/x/lottery/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.TxCounter != nil {
		k.SetTxCounter(ctx, *genState.TxCounter)
	}
	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, elem)
	}

	// Set user count
	k.SetUserCount(ctx, genState.UserCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all txCounter
	txCounter, found := k.GetTxCounter(ctx)
	if found {
		genesis.TxCounter = &txCounter
	}
	genesis.UserList = k.GetAllUser(ctx)
	genesis.UserCount = k.GetUserCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
