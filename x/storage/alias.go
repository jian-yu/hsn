package storage

import (
	"github.com/shegaoyuan/hsn/x/storage/internal/keeper"
	"github.com/shegaoyuan/hsn/x/storage/internal/types"
)

const (
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	DefaultBondDenom  = types.DefaultBondDenom
)

var (
	NewKeeper           = keeper.NewKeeper
	NewStorageProvider  = types.NewStorageProvider
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewParams           = types.NewParams
	ValidateParams      = types.ValidateParams
	DefaultParams       = types.DefaultParams
	NewQuerier          = keeper.NewQuerier

	ModuleCdc          = types.ModuleCdc
	StorageProviderKey = types.StorageProviderKey
	KeyStorageDenom    = types.KeyStorageDenom
	KeyCostPerByte     = types.KeyCostPerByte
	KeyCapatity        = types.KeyCapatity
	KeyLength          = types.KeyLength
)

type (
	Keeper          = keeper.Keeper
	GenesisState    = types.GenesisState
	StorageProvider = types.StorageProvider
	Params          = types.Params
)
