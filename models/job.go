package models

// Job models a deployment of containers to perform a time-boxed task
type Job struct {
	// Job name
	Name string `yaml:"name"`
	// ImageName is the container image name
	ImageName string `yaml:"image"`
	// ImageTag is the container image tag
	ImageTag string `yaml:"tag"`
	// Command is the command trigger for the job
	Command string `yaml:"command"`
	// Requests specify resource requests
	Requests ResourceSpec `yaml:"requests,omitempty"`
	// Limits specify resource limits
	Limits ResourceSpec `yaml:"limits,omitempty"`
	// NodeAffinity contains list of key value pairs to match to nodes
	NodeAffinity []KeyValue `yaml:"nodeAffinity,omitempty"`
	// Tolerations contain a list of tolerations
	Tolerations []KeyValue `yaml:"tolerations,omitempty"`
	// Job restart policy
	RestartPolicy string `yaml:"restartPolicy"`
	// Backoff limit of job
	BackoffLimit string `yaml:"backoffLimit"`
}
