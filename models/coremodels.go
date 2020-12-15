package models

// Port is a service or container port descriptor
type Port struct {
	// Name is the port name
	Name       string `yaml:"name,omitempty"`
	TargetPort int    `yaml:"targetPort"`
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

// // ChartValues is a structure representing input parameters of a chart
// type ChartValues struct {
// 	Image struct {
// 		Registry   string `yaml:"registry"`
// 		Repository string `yaml:"repository"`
// 		Tag        string `yaml:"tag"`
// 		PullPolicy string `yaml:"pullPolicy"`
// 	} `yaml:"image"`

// 	Service struct {

// 	}
// }
