package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

// JobTemplate is a base template for helm chart jobs
const JobTemplate = `
{{- $outer := . }}
{{- if .Application.Spec.Jobs }}
{{- range .Application.Spec.Jobs }}
{{- range $key,$value := . }}
---
apiVersion: batch/v1
kind: Job
metadata:
  name: {{ $key }}
  labels:
    app: {{ $key }}
spec:
  backoffLimit: {{"{{"}} .Values.{{ $key }}.backoffLimit {{"}}"}}
  template:
    metadata:
      labels:
        app: {{ $key }}
    spec:
      restartPolicy: Never
      containers:
      - name: {{ $key }}
        image: {{"{{"}} .Values.{{ $key }}.image.repository {{"}}"}}:{{"{{"}} .Values.{{ $key }}.image.tag {{"}}"}}
        imagePullPolicy: {{"{{"}} .Values.pullPolicy {{"}}"}}
		{{- if or ($value.ConfigMaps) ($value.Secrets) }}
        env:
		{{- if $value.ConfigMaps }}
		{{- range $value.ConfigMaps }}
		   {{- $cmap := . }}
		   {{- range $outer.Application.Spec.ConfigMaps }}
		   {{- range $key,$value := . }}
		   {{- if eq ($key) ($cmap) }}
			 {{- range $value.Data }}
			   {{- range $key,$value := . }}
        - name: {{ $key }}
          valueFrom:
            configMapKeyRef:
              name: {{ $cmap }}
              key: {{ $key }}
			   {{- end }}
			 {{- end }}
		   {{- end }}
		   {{- end }}
		   {{- end }}
		{{- end }}
		{{- end }}
		{{- if $value.Secrets }}
		{{- range $value.Secrets }}
		   {{- $secret := . }}
		   {{- range $outer.Application.Spec.Secrets }}
		   {{- range $key,$value := . }}
		   {{- if eq ($key) ($secret) }}
			 {{- range $value.Data }}
        - name: {{ . }}
          valueFrom:
            secretKeyRef:
              name: {{ $secret }}
              key: {{ . }}
			 {{- end }}
		   {{- end }}
		   {{- end }}
		   {{- end }}
		{{- end }}
		{{- end }}
		{{- end }}
        resources:
          requests: {{"{{"}} .Values.{{ $key }}.resources.requests {{"}}"}}
          limits: {{"{{"}} .Values.{{ $key }}.resources.limits {{"}}"}}
		{{- if $value.Storage }}
        volumeMounts:
		{{- range $value.Storage }}
		{{- range $key,$value := . }}
          - name: {{ $key }}-data
            mountPath: {{ $value.Mount }}
		{{- end }}
		{{- end }}
		{{- end }}
	  {{- if $value.Storage }}
      volumes:
	  {{- range $value.Storage }}
	  {{- range $key,$value := . }}
        - name: {{ $key }}-data
          persistentVolumeClaim: 
            claimName: {{ $key }}
	  {{- end }}
	  {{- end }}
	  {{- end }}
{{- end }}
{{- end }}
{{- end }}
`

// WriteJobs outputs the deployment templates for these charts
func WriteJobs(metaApp models.MetaApplication, outDir string) {
	utils.OutputTemplate(metaApp, JobTemplate, filepath.Join(outDir, "jobs.yaml"))
}
