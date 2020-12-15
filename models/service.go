package models

// Service models a service endpoint
type Service struct {
	// Service name
	Name string
	//
	AppName string
	// Ports is a list of port descriptors
	Ports []Port
}
