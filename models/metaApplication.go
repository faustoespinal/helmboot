package models

// MetaApplication defines full template parameters for templatized helm resources.
type MetaApplication struct {
	Meta struct {
		ReleaseName string `yaml:"releaseName"`
		Namespace   string `yaml:"namespace"`
	} `yaml:"meta"`
	Application Application `yaml:"application"`
}

// CreateMetaApplication initializes an instance of a meta application
func CreateMetaApplication(app Application) MetaApplication {
	var metaApp = MetaApplication{}

	metaApp.Meta.ReleaseName = "{{ .ReleaseName }}"
	metaApp.Meta.Namespace = "{{ .Release.Namespace }}"
	metaApp.Application = app
	return metaApp
}
