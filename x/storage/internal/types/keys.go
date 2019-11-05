package types

var StorageProviderKey = []byte{0x11}

const (
	// module name
	ModuleName = "storage"
	// default paramspace for params keeper
	DefaultParamspace = ModuleName
	// StoreKey is the default store key for storage
	StoreKey = ModuleName
	// QuerierRoute is the querier route for the storage store.
	QuerierRoute = StoreKey
	// Default denom
	DefaultBondDenom = "uhsn"
	QueryParameters  = "parameters"
)
