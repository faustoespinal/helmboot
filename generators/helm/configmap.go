package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// ConfigMapTemplate defines a template of a kubernetes configmap
const ConfigMapTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.ConfigMaps }}
{{- range .Application.Spec.ConfigMaps }}
{{- range $key,$value := . }}
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ $key }}
data:
  {{- range $value.Data }}
  {{- range $key,$value := . }}
  {{ $key }}: {{ $value }}
  {{- end }}
  {{- end }}
{{- end }}
{{- end }}
{{- end }}
`

// WriteConfigmaps outputs the config-map templates for these charts
func WriteConfigmaps(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, ConfigMapTemplate, filepath.Join(outDir, "configmaps.yaml"))
}
