package models

// Secret encodes application private information items
type Secret struct {
	Type string   `yaml:"type"`
	Data []string `yaml:"data"`
}
