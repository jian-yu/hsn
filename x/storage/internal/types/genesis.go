package types

type GenesisState struct {
	StorageProvider StorageProvider `json:"storageProvider" yaml:"storageProvider"`
	Params          Params          `json:"params" yaml:"params"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(storageProvider StorageProvider, params Params) GenesisState {
	return GenesisState{
		StorageProvider: storageProvider,
		Params:          params,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{
		StorageProvider: DefaultInitialStorageProvider(),
		Params:          DefaultParams(),
	}
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	err := ValidateParams(data.Params)
	if err != nil {
		return err
	}
	return ValidateStorageProvider(data.StorageProvider)
}
