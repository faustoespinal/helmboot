package models

// Pvc is a storage claim made by a cloud-native application
type Pvc struct {
	// Name is the storage claim name
	Name string `yaml:"name"`
	// StorageClass is the type of storage desired
	StorageClass string `yaml:"storageClass,omitempty"`
	// StorageRequest is the amount of storage requested
	StorageRequest string `yaml:"storageRequest,omitempty"`
}
