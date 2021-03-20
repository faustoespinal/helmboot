package edison

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// EpaTemplate generates templates for an edison postgres database connection resource
const EpaTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.Databases }}
{{- range .Application.Spec.Databases }}
---
apiVersion: ees.ge.com/v1
kind: EesPostgresAccount
metadata:
  annotations:
    resource/author: {{ $outer.Application.Name }}
  labels:
    targetHost: {{"{{"}} .Values.epa.{{ regexReplaceAll "\\W+" . "_" }}.postgres.targetname {{"}}"}}
  name: {{ . }}
spec:
  clientid: {{ $outer.Application.Name }}
  clientns: {{ $outer.Meta.Namespace }}
  dbname: {{"{{"}} .Values.epa.{{ regexReplaceAll "\\W+" . "_" }}.postgres.dbname {{"}}"}}
  targetname: {{"{{"}} .Values.epa.{{ regexReplaceAll "\\W+" . "_" }}.postgres.targetname {{"}}"}}
  username: {{"{{"}} .Values.epa.{{ regexReplaceAll "\\W+" . "_" }}.postgres.username {{"}}"}}
{{- end }}
{{- end }}
`

// WriteEpa outputs the database connection templates for these charts
func WriteEpa(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, EpaTemplate, filepath.Join(outDir, "epa.yaml"))
}
