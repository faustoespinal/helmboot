package models

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
