package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// PvcTemplate defines the base template for a Persistent volume claim resource
const PvcTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.Storage }}
{{- range .Application.Spec.Storage }}
{{- range $key,$value := . }}
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{ $key }}
spec:
  accessModes:
	- {{ $value.Mode }}
  {{- if $value.StorageClass }}
  storageClassName: {{ $value.StorageClass }}
  {{- end }}
  resources:
    requests:
      storage: {{"{{"}} .Values.{{ $key }}.size {{"}}"}}
{{- end}}
{{- end}}
{{- end}}
`

// WritePvcs outputs the persistent volume claim templates for the charts
func WritePvcs(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, PvcTemplate, filepath.Join(outDir, "pvcs.yaml"))
}
