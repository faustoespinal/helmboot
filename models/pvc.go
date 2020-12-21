package models

// Pvc describes a persistent storage request for an application
type Pvc struct {
	Size         string `yaml:"size"`
	Mode         string `yaml:"mode,omitempty"`
	StorageClass string `yaml:"storageClass,omitempty"`
}
