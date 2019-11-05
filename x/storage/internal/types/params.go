package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

var (
	KeyStorageDenom = []byte("StorageDenom")
	KeyCostPerByte  = []byte("CostPerByte")
	KeyCapatity     = []byte("Capatity")
	KeyLength       = []byte("Length")
)

type Params struct {
	StorageDenom string   `json:"storage_denom" yaml:"storage_denom"` // type of coin for storage
	CostPerByte  sdk.Coin `json:"cost_per_byte" yaml:"cost_per_byte"` //the cost of per byte
	Capatity     sdk.Int  `json:"capatity" yaml:"capatity"`           //the capatity of storageProvider
	Length       sdk.Int  `json:"length" yaml:"length"`               //the length of storageProvider's used capatity
}

// ParamTable for storage module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(denom string, costPerByte sdk.Coin, capatity sdk.Int, length sdk.Int) Params {
	return Params{
		StorageDenom: denom,
		CostPerByte:  costPerByte,
		Capatity:     capatity,
		Length:       length,
	}
}

// validate params
func ValidateParams(params Params) error {
	if params.StorageDenom == "" {
		return fmt.Errorf("storage parameter StorageDenom can't be an empty string")
	}
	if !params.CostPerByte.IsValid() {
		return fmt.Errorf("storage parameter CostPerByte is invalid")
	}
	if params.Capatity.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("storage parameter Capatity  can't be less than or equal to zero")
	}
	if params.Length.LT(sdk.ZeroInt()) {
		return fmt.Errorf("storage parameter Capatity  can't be less than zero")
	}
	return nil
}

// default storage module parameters
func DefaultParams() Params {
	return Params{
		StorageDenom: sdk.DefaultBondDenom,
		CostPerByte:  sdk.NewCoin(DefaultBondDenom, sdk.NewInt(1)),
		Capatity:     sdk.NewInt(1000 * 1000 * 1024), // default capatity is 10GB
		Length:       sdk.NewInt(0),
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyStorageDenom, Value: &p.StorageDenom},
		{Key: KeyCostPerByte, Value: &p.CostPerByte},
		{Key: KeyCapatity, Value: &p.Capatity},
		{Key: KeyLength, Value: &p.Length},
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`Storage Params:
  Storage Denom:             %s
  CostPerByte: 				 %s
  Capatity:					 %s
  Length:					 %s
`,
		p.StorageDenom, p.CostPerByte, p.Capatity, p.Length,
	)
}
