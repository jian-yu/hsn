package storage

import (
	"encoding/json"
	"math/big"

	"github.com/gogo/protobuf/codec"
	"github.com/gorilla/mux"
	"github.com/shegaoyuan/hsn/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var _ module.AppModuleBasic = AppMoudleBasic{}
var _ module.AppModule = AppModule{}

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := types.ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// Once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(data)
}

func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	//rpc.RegisterRoutes(ctx, rtr, StoreKey)
}

func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {

	// return cli.GetQueryCmd(types.ModuleName, cdc)
	return nil
}

func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {

	// return cli.GetTxCmd(types.ModuleName, cdc)
	return nil
}

type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

func NewAppModule(keeper Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

func (AppModule) Name() string {
	return types.ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string {
	return types.RouterKey
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

func (am AppModule) QuerierRoute() string {
	return types.ModuleName
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

func (am AppModule) BeginBlock(ctx sdk.Context, bl abci.RequestBeginBlock) {
	// Consider removing this when using evm as module without web3 API
	bloom := ethtypes.BytesToBloom(am.keeper.bloom.Bytes())
	am.keeper.SetBlockBloomMapping(ctx, bloom, bl.Header.GetHeight()-1)
	am.keeper.SetBlockHashMapping(ctx, bl.Header.LastBlockId.GetHash(), bl.Header.GetHeight()-1)
	am.keeper.bloom = big.NewInt(0)
	am.keeper.txCount.reset()
}

func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	// Gas costs are handled within msg handler so costs should be ignored
	ebCtx := ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())

	// Commit state objects to KV store
	_, err := am.keeper.csdb.WithContext(ebCtx).Commit(true)
	if err != nil {
		panic(err)
	}

	return []abci.ValidatorUpdate{}
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return InitGenesis(ctx, am.keeper, genesisState)
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return types.ModuleCdc.MustMarshalJSON(gs)
}
