package edison

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// EraTemplate generates templates for an edison amqp messaging connection resource
const EraTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.Messaging }}
{{- range .Application.Spec.Messaging }}
---
apiVersion: ees.ge.com/v1
kind: EesRabbitmqAccount
metadata:
  annotations:
    resource/author: {{ $outer.Application.Name }}
  name: {{ . }}
spec:
  clientid: {{ $outer.Application.Name }}
  clientns: {{ $outer.Meta.Namespace }}
  username: {{"{{"}} .Values.era.{{ regexReplaceAll "\\W+" . "_" }}.amqp.username {{"}}"}}
  vhostname: {{"{{"}} .Values.era.{{ regexReplaceAll "\\W+" . "_" }}.amqp.vhostname {{"}}"}}
{{- end }}
{{- end }}
`

// WriteEra outputs the messaging templates for these charts
func WriteEra(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, EraTemplate, filepath.Join(outDir, "era.yaml"))
}
