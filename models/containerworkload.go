package models

// ContainerWorkload describes a workload run as a deployment or job
type ContainerWorkload struct {
	// Image base name
	Image string `yaml:"image"`
	// Version tag of the image
	Tag        string   `yaml:"tag"`
	ConfigMaps []string `yaml:"configmaps,omitempty"`
	Secrets    []string `yaml:"secrets,omitempty"`
	Port       int      `yaml:"port"`
	// List of storage mounts
	Storage []map[string]StorageMount `yaml:"storage,omitempty"`
	// List of database connections
	Databases []string `yaml:"databases,omitempty"`
	// List of messaging connections
	Messaging []string `yaml:"messaging,omitempty"`
	Resources *struct {
		Requests *ResourceSpec `yaml:"requests,omitempty"`
		Limits   *ResourceSpec `yaml:"limits,omitempty"`
	} `yaml:"resources,omitempty"`
}
