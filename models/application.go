package models

// Application models the input for generating a cloud native application deployment
type Application struct {
	// ApiVersion is the version of the input schema
	APIVersion string `yaml:"apiVersion"`
	// Type is the category of the application: application, service, job
	Type        string `yaml:"type"`
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
	Version     string `yaml:"version"`
	AppVersion  string `yaml:"appVersion,omitempty"`
	Spec        struct {
		Security struct {
			GrantTypes []string             `yaml:"grantTypes"`
			Roles      []map[string]AppRole `yaml:"roles"`
		} `yaml:"security"`
		// List of deployments
		Deployments []map[string]Workload `yaml:"deployments,omitempty"`
		// List of time-boxed jobs
		Jobs       []map[string]Workload     `yaml:"jobs,omitempty"`
		Services   []map[string]AppService   `yaml:"services,omitempty"`
		Ingresses  []map[string]AppIngress   `yaml:"ingresses,omitempty"`
		ConfigMaps []map[string]AppConfigMap `yaml:"configmaps,omitempty"`
		Secrets    []map[string]AppSecret    `yaml:"secrets,omitempty"`
		Storage    []map[string]AppStorage   `yaml:"storage,omitempty"`
		Databases  []string                  `yaml:"databases,omitempty"`
		Messaging  []string                  `yaml:"messaging,omitempty"`
	} `yaml:"spec"`
}

// AppRole describes an application specific role
type AppRole struct {
	Scopes []string `yaml:"scopes"`
}

// AppService is a descriptor for an application service
type AppService struct {
	Deployment string `yaml:"deployment"`
}

// AppStorage describes a unit of storage for an application
type AppStorage struct {
	Size         string `yaml:"size"`
	Mode         string `yaml:"mode,omitempty"`
	StorageClass string `yaml:"storageClass,omitempty"`
}

// AppIngress describes an ingress
type AppIngress struct {
	Service         string `yaml:"service"`
	Namespace       string `yaml:"namespace,omitempty"`
	ExternalService string `yaml:"externalService,omitempty"`
}

// AppSecret encodes application private information items
type AppSecret struct {
	Type string   `yaml:"type"`
	Data []string `yaml:"data"`
}

// AppConfigMap defines alternate ways of encoding configuration data in a file and/or key/value pairs
type AppConfigMap struct {
	Data []map[string]string `yaml:"data,omitempty"`
}

// StorageMount describes the way to mount a unit of storage
type StorageMount struct {
	Mount string `yaml:"mount,omitempty"`
}

// Workload describes a workload run as a deployment or job
type Workload struct {
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
}
