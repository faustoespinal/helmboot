package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// SecretTemplate defines a template of a kubernetes configmap
const SecretTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.Secrets }}
{{- range .Application.Spec.Secrets }}
{{- range $key,$value := . }}
---
apiVersion: v1
kind: Secret
metadata:
  name: {{ $key }}
data:
  {{- range $value.Data }}
  {{ . }}: {{"{{"}} .Values.secret.{{ $key }}.{{ . }} | b64enc {{"}}"}}
  {{- end }}
{{- end }}
{{- end }}
{{- end }}
`

// WriteSecrets outputs the secret templates for these charts
func WriteSecrets(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, SecretTemplate, filepath.Join(outDir, "secrets.yaml"))
}
