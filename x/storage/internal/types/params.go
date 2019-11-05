package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

var (
	KeyStorageDenom = []byte("StorageDenom")
)

type Params struct {
	StorageDenom string `json:"storage_denom" yaml:"storage_denom"` // type of coin for storage
}

// ParamTable for storage module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(denom string) Params {
	return Params{
		StorageDenom: denom,
	}
}

// validate params
func ValidateParams(params Params) error {
	if params.StorageDenom == "" {
		return fmt.Errorf("storage parameter StorageDenom can't be an empty string")
	}
	return nil
}

// default storage module parameters
func DefaultParams() Params {
	return Params{
		StorageDenom: sdk.DefaultBondDenom,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyStorageDenom, Value: &p.StorageDenom},
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`Storage Params:
  Storage Denom:             %s

`,
		p.StorageDenom,
	)
}
