package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

const appRegistrationTmpl = `
apiVersion: ees.ge.com/v1
kind: ApplicationRegistration
metadata:
  name: {{.Application.Name}}
spec:
  accessTokenDuration: 3600
  appName: {{.Application.Name}}
  callbackUrls:
  - https://<hostname>/image-application-execution-user/
  - https://<hostname>/image-application-execution-login
  claimDialects: []
  description: {{.Application.Name}}
  displayName: {{.Application.Description}}
  grantTypes:
  {{- range .Application.Spec.Security.GrantTypes }}
  - {{ . }}
  {{- end }}
  refreshTokenDuration: 86400
  roles:
  {{- range .Application.Spec.Security.Roles }}
  {{- range $key, $value := . }}
  - roleName: {{ $key }}
    scopes:
	{{- range $value.Scopes }}
    - {{ . }}
	{{- end }}
  {{- end }}
  {{- end }}
  tokenIssuer: JWT
`

type valueAppReg struct {
	Application models.Application
}

// WriteApplicationRegistration outputs a template for an app registration resource
func WriteApplicationRegistration(application models.Application, outDir string) {
	content := valueAppReg{}
	content.Application = application

	utils.OutputTemplate(content, appRegistrationTmpl, filepath.Join(outDir, "app-registration.yaml"))
}
