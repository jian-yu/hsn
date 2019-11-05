package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/shegaoyuan/hsn/x/storage/internal/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc           *codec.Codec
	storeKey      sdk.StoreKey
	paramSpace    params.Subspace
	stakingKeeper types.StakingKeeper
	supplyKeeper  types.SupplyKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace, stakingKeeper types.StakingKeeper, supplyKeeper types.SupplyKeeper) Keeper {
	if addr := supplyKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the storage module account has not been set")
	}
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		paramSpace:    paramSpace,
		stakingKeeper: stakingKeeper,
		supplyKeeper:  supplyKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetStorageProvider(ctx sdk.Context) (storageProvider types.StorageProvider) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.StorageProviderKey)
	if b == nil {
		panic("stored StorageProvider should not have been nil")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &storageProvider)
	return
}

//SetStorageProvider set the storageprovider
func (k Keeper) SetStorageProvider(ctx sdk.Context, storageProvider types.StorageProvider) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(storageProvider)
	store.Set(types.StorageProviderKey, b)
}

// GetParams returns the total set of storage parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of storage parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
