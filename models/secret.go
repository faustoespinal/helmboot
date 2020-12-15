package models

// Secret stores private information
type Secret struct {
	// Name is the name of the secret
	Name string `yaml:"name"`
	// SecretFiles contains list of references of files to encode as secrets
	SecretFiles []struct {
		FileName string `yaml:"fileName"`
		FilePath string `yaml:"filePath"`
	} `yaml:"secretFiles"`
}
