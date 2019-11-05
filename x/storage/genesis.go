package storage

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new stroage genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.SetStorageProvider(ctx, data.StorageProvider)
	keeper.SetParams(ctx, data.Params)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	storageProvider := keeper.GetStorageProvider(ctx)
	params := keeper.GetParams(ctx)
	return NewGenesisState(storageProvider, params)
}
