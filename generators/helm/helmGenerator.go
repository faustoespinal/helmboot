package helm

import (
	edison "helmboot/generators/helm/edison"
	"helmboot/models"
	"path/filepath"

	"go.uber.org/zap"
)

// Generator is an implementation of the Templator interface
type Generator struct {
}

// Name returns descriptive name of the template generator
func (g *Generator) Name() string {
	return "Helm"
}

// Write the templates to the specified output directory
func (g *Generator) Write(application models.Application, outDir string) {
	zap.S().Infof("Generating %s to directory: %s\n", application.Name, outDir)

	//jsonString, _ := utils.PrettyJSON(application)
	//zap.S().Infof("Application: %v\n", jsonString)

	WriteHelmBase(application, outDir)
	WriteReadmeMd(application, outDir)
	WriteValues(application, outDir)

	metaApp := models.CreateMetaApplication(application)
	templateDir := filepath.Join(outDir, "templates")
	WriteApplicationRegistration(application, templateDir)

	if len(application.Spec.Services) > 0 {
		WriteServices(metaApp, templateDir)
	}
	if len(application.Spec.Ingresses) > 0 {
		WriteIngresses(metaApp, templateDir)
	}
	if len(application.Spec.Deployments) > 0 {
		WriteDeployments(metaApp, templateDir)
	}
	if len(application.Spec.Jobs) > 0 {
		WriteJobs(metaApp, templateDir)
	}
	if len(application.Spec.Deployments) > 0 || len(application.Spec.Jobs) > 0 {
		WriteServiceAccounts(metaApp, templateDir)
	}
	if len(application.Spec.ConfigMaps) > 0 {
		WriteConfigmaps(metaApp, templateDir)
	}
	if len(application.Spec.Secrets) > 0 {
		WriteSecrets(metaApp, templateDir)
	}
	if len(application.Spec.Storage) > 0 {
		WritePvcs(metaApp, templateDir)
	}

	isEdison := true
	if isEdison {
		if len(application.Spec.Messaging) > 0 {
			edison.WriteEra(metaApp, templateDir)
		}
		if len(application.Spec.Databases) > 0 {
			edison.WriteEpa(metaApp, templateDir)
		}
	}
	if application.Spec.Testing != nil {
		WriteSvcTests(metaApp, filepath.Join(templateDir, "tests"))
	}
}
