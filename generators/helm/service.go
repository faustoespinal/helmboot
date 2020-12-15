package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// SvcTemplate defines a template of a kubernetes service
const SvcTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.Services }}
{{- range .Application.Spec.Services }}
{{- range $key,$value := . }}
---
apiVersion: v1
kind: Service
metadata:
  name: {{ $key }}
spec:
  type: {{"{{"}} .Values.service.{{ $key }}.type {{"}}"}}
  ports:
    - name: {{ $key }}
      port: {{"{{"}} .Values.service.{{ $key }}.port {{"}}"}}
      targetPort: {{ $value.Deployment }}
  selector:
    app: {{ $value.Deployment }}
{{- end }}
{{- end }}
{{- end }}
`

// WriteServices outputs the service templates for these charts
func WriteServices(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, SvcTemplate, filepath.Join(outDir, "service.yaml"))
}
