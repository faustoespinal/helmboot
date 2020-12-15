package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

const gitIgnore = `
# Patterns to ignore when building packages.
# This supports shell glob matching, relative path matching, and
# negation (prefixed with !). Only one pattern per line.
.DS_Store
# Common VCS dirs
.git/
.gitignore
.bzr/
.bzrignore
.hg/
.hgignore
.svn/
# Common backup files
*.swp
*.bak
*.tmp
*.orig
*~
# Various IDEs
.project
.idea/
*.tmproj
.vscode/
`

const notesTxt = `
Workload {{.Name}} has been deployed
`

const chartYaml = `
apiVersion: v2
name: {{.Name}}
description: {{.Description}}
type: application

# This is the chart version. This version number should be incremented each time you make changes
# to the chart and its templates, including the app version.
# Versions are expected to follow Semantic Versioning (https://semver.org/)
version: {{.Version}}

# This is the version number of the application being deployed. This version number should be
# incremented each time you make changes to the application. Versions are not expected to
# follow Semantic Versioning. They should reflect the version the application is using.
appVersion: {{.AppVersion}}
`

// WriteHelmBase writes the initial skaffolding for a helm chart using the application model information
func WriteHelmBase(application models.Application, outDir string) {
	templatesDir := filepath.Join(outDir, "templates")
	testDir := filepath.Join(templatesDir, "tests")
	utils.CreateDir(testDir)

	err := utils.WriteFile(gitIgnore, filepath.Join(outDir, ".gitignore"))
	if err != nil {
		panic(err)
	}

	chartYamlPath := filepath.Join(outDir, "Chart.yaml")
	utils.OutputTemplate(application, chartYaml, chartYamlPath)

	notesPath := filepath.Join(templatesDir, "NOTES.txt")
	utils.OutputTemplate(application, notesTxt, notesPath)
}
