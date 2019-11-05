package types

type StorageProvider struct {
}

func NewStorageProvider() StorageProvider {
	return StorageProvider{}
}

func DefaultInitialStorageProvider() StorageProvider {
	return StorageProvider{}
}

// validate storageProvider
func ValidateStorageProvider(storageProvider StorageProvider) error {

	return nil
}
