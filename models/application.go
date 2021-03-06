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
	IsEdison    bool
	Spec        struct {
		Security struct {
			GrantTypes []string             `yaml:"grantTypes"`
			Roles      []map[string]AppRole `yaml:"roles"`
		} `yaml:"security,omitempty"`
		// List of deployments
		Deployments []map[string]ContainerWorkload `yaml:"deployments,omitempty"`
		// List of time-boxed jobs
		Jobs       []map[string]ContainerWorkload `yaml:"jobs,omitempty"`
		Services   []map[string]AppService        `yaml:"services,omitempty"`
		Ingresses  []map[string]AppIngress        `yaml:"ingresses,omitempty"`
		ConfigMaps []map[string]AppConfigMap      `yaml:"configmaps,omitempty"`
		Secrets    []map[string]Secret            `yaml:"secrets,omitempty"`
		Storage    []map[string]Pvc               `yaml:"storage,omitempty"`
		Databases  []string                       `yaml:"databases,omitempty"`
		Messaging  []string                       `yaml:"messaging,omitempty"`
		Testing    *struct {
			Image   string   `yaml:"image"`
			Command []string `yaml:"command"`
			Args    []string `yaml:"args,omitempty"`
		} `yaml:"testing,omitempty"`
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
