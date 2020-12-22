package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// ServiceAccountTemplate defines a template of a kubernetes configmap
const ServiceAccountTemplate = `
{{"{{-"}} if .Values.serviceAccount.create {{"}}"}}
{{- $outer := . }}
{{- if .Application.Spec.Deployments }}
{{- range .Application.Spec.Deployments }}
{{- range $key,$value := . }}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ $key }}
  labels:
    account: {{ $key }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Application.Spec.Jobs }}
{{- range .Application.Spec.Jobs }}
{{- range $key,$value := . }}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ $key }}
  labels:
    account: {{ $key }}
{{- end }}
{{- end }}
{{- end }}

{{"{{-"}} end {{"}}"}}
`

// WriteServiceAccounts outputs the service account templates for these charts
func WriteServiceAccounts(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, ServiceAccountTemplate, filepath.Join(outDir, "serviceaccounts.yaml"))
}
