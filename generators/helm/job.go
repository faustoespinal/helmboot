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
    sourceapp: {{ $outer.Application.Name }}
    sourceversion: {{ $outer.Application.Version }}
    sourceappversion: {{ $outer.Application.AppVersion }}
spec:
  backoffLimit: {{"{{"}} .Values.{{ $key }}.backoffLimit {{"}}"}}
  template:
    metadata:
      labels:
        app: {{ $key }}
    spec:
      serviceAccountName: {{ $key }}
      restartPolicy: Never
      containers:
      - name: {{ $key }}
        image: {{"{{"}} .Values.{{ $key }}.image.repository {{"}}"}}:{{"{{"}} .Values.{{ $key }}.image.tag {{"}}"}}
        imagePullPolicy: {{"{{"}} .Values.pullPolicy {{"}}"}}
		{{- if or ($value.Env) ($value.ConfigMaps) ($value.Secrets) ($value.Databases) }}
        env:
    {{- if $value.Env }}
    {{- range $value.Env }}
        - name: {{ .Name }}
          value: "{{ .Value }}"
    {{- end }}
    {{- end }}    
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
		{{- if $value.Databases }}
		{{- range $value.Databases }}
        - name: EIS_DB_USER
          valueFrom:
            secretKeyRef:
              name: {{ . }}-eespostgresaccount-secret
              key: user
        - name: EIS_DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: {{ . }}-eespostgresaccount-secret
              key: password
        - name: EIS_DB_IP
          valueFrom:
            configMapKeyRef:
              name: {{ . }}-eespostgresaccount-configmap
              key: service-host
        - name: EIS_DB_PORT
          valueFrom:
            configMapKeyRef:
              name: {{ . }}-eespostgresaccount-configmap
              key: service-port
        - name: EIS_DB_NAME
          valueFrom:
            configMapKeyRef:
              name: {{ . }}-eespostgresaccount-configmap
              key: dbname		
        {{- end }}
        {{- end }}
        {{- if $value.Messaging }}
        {{- range $value.Messaging }}
            - name: EIS_RABBITMQ_HOST
              valueFrom:
                configMapKeyRef:
                  name: {{ . }}-eesrabbitmqaccount-configmap
                  key: rabbitmq-service-host
            - name: EIS_RABBITMQ_PORT
              valueFrom:
                configMapKeyRef:
                  name: {{ . }}-eesrabbitmqaccount-configmap
                  key: rabbitmq-service-port
            - name: EIS_RABBITMQ_VHOST
              valueFrom:
                configMapKeyRef:
                  name: {{ . }}-eesrabbitmqaccount-configmap
                  key: rabbitmq-service-vhost
            - name: EIS_RABBITMQ_USERNAME
              valueFrom:
                secretKeyRef:
                  name: {{ . }}-eesrabbitmqaccount-secret
                  key: rabbitmq-user
            - name: EIS_RABBITMQ_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: {{ . }}-eesrabbitmqaccount-secret
                  key: rabbitmq-password		
        {{- end }}
        {{- end }}
        {{- end }}
        securityContext:
          runAsUser: 1000
        resources: {{"{{"}} .Values.{{ $key }}.resources {{"}}"}}
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
