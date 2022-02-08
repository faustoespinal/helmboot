package models

// ChartValue describes a helm chart value
type ChartValue struct {
	Name         string
	Description  string
	DefaultValue string
}

// Port is a service or container port descriptor
type Port struct {
	// Name is the port name
	Name string `yaml:"name,omitempty"`
	// TargetPort is the destination port
	TargetPort int `yaml:"targetPort"`
}

// KeyValue is a simple model for key/value pair
type KeyValue struct {
	// Key is the key part of tuple
	Key string
	// Value is the value part of tuple
	Value string
}

// ResourceSpec is a limit specification for CPU and memory
type ResourceSpec struct {
	CPU    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}

// AppRole describes an application specific role
type AppRole struct {
	Scopes []string `yaml:"scopes"`
}

// AppService is a descriptor for an application service
type AppService struct {
	Deployment string `yaml:"deployment"`
}

// AppIngress describes an ingress
type AppIngress struct {
	Service         string `yaml:"service"`
	Namespace       string `yaml:"namespace,omitempty"`
	ExternalService string `yaml:"externalService,omitempty"`
}

// AppConfigMap defines alternate ways of encoding configuration data in a file and/or key/value pairs
type AppConfigMap struct {
	Data []map[string]string `yaml:"data,omitempty"`
}

// StorageMount describes the way to mount a unit of storage
type StorageMount struct {
	Mount string `yaml:"mount,omitempty"`
}
