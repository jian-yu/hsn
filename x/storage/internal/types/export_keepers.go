package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StakingKeeper interface {
}

type SupplyKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
}
