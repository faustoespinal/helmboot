package models

// Deployment models a deployment of containers as a long lasting workload
type Deployment struct {
	// Deployment name
	Name string `yaml:"name"`
	// ImageName is the container image name
	ImageName string `yaml:"imageName"`
	// ImageTag is the container image tag
	ImageTag string `yaml:"imageTag"`
	// Replicas contains the number of replica containers in this deployment.
	Replicas int `yaml:"replicas"`
	// Ports is a list of port descriptors
	Ports []Port `yaml:"ports"`
	// Requests specify resource requests
	Requests ResourceSpec `yaml:"requests,omitempty"`
	// Limits specify resource limits
	Limits ResourceSpec `yaml:"limits,omitempty"`
	// NodeAffinity contains list of key value pairs to match to nodes
	NodeAffinity []KeyValue `yaml:"nodeAffinity,omitempty"`
	// Tolerations contain a list of tolerations
	Tolerations []KeyValue `yaml:"tolerations,omitempty"`
}
