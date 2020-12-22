package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// IngressTemplate defines a template of a kubernetes service
const IngressTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.Ingresses }}
{{- range .Application.Spec.Ingresses }}
{{- range $key,$value := . }}
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    plugins.konghq.com: jwt,authz
  name: {{ $key }}
spec:
  rules:
  - http:
      paths:
        - path: /{{ $outer.Application.Name }}/{{ $value.Service }}
          pathType: Prefix
          backend:
            service:
              name: {{ $value.Service }}
              port:
                number: {{"{{"}} .Values.service.{{ snakecase $value.Service }}.port {{"}}"}}
---
apiVersion: configuration.konghq.com/v1
kind: KongIngress
metadata:
  name: {{ $key }}-kong
proxy:
  path: /{{ $outer.Application.Name }}/{{ $value.Service }}
route:
  strip_path: true
  preserve_host: false
{{- if $value.Namespace }}
---
apiVersion: v1
kind: Service
metadata:
  name: {{ $value.Service }}
spec:
  type: ExternalName
  externalName: {{ $value.ExternalService }}.{{ $value.Namespace }}.cluster.local
{{- end }}
{{- end }}
{{- end }}
{{- end }}
`

// WriteIngresses outputs the ingress templates for these charts
func WriteIngresses(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, IngressTemplate, filepath.Join(outDir, "ingress.yaml"))
}
