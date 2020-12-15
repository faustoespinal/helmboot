package helm

import (
	"helmboot/models"
	"path/filepath"

	"github.com/golang/glog"
)

// Generator is an implementation of the Templator interface
type Generator struct {
}

// Name returns descriptive name of the template generator
func (g *Generator) Name() string {
	return "HelmGenerator"
}

// Write the templates to the specified output directory
func (g *Generator) Write(application models.Application, outDir string) {
	glog.Infof("Generating %s to directory: %s\n", application.Name, outDir)

	glog.Infof("Application: %#v\n", application)
	WriteHelmBase(application, outDir)
	WriteReadmeMd(application, outDir)
	WriteValues(application, outDir)

	metaApp := models.CreateMetaApplication(application)
	templateDir := filepath.Join(outDir, "templates")
	WriteApplicationRegistration(application, templateDir)
	WriteServices(metaApp, templateDir)
	WriteIngresses(metaApp, templateDir)
	WriteDeployments(metaApp, templateDir)
	WriteJobs(metaApp, templateDir)
	WriteConfigmaps(metaApp, templateDir)
	WriteSecrets(metaApp, templateDir)

	WritePvcs(metaApp, templateDir)
}
