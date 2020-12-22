package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"

	"github.com/golang/glog"
)

const workloadTmpl = `
  ## Affinity for pod assignment
  ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
  ## Note: podAffinityPreset, podAntiAffinityPreset, and  nodeAffinityPreset will be ignored when it's set
  ##
  affinity: {}

  ## Node labels for pod assignment
  ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
  ##
  nodeSelector: {}

  ## Tolerations for pod assignment
  ## Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
  ##
  tolerations: []
`

const valuesTmpl = `
registry: docker.io
{{- if .Spec.Deployments}}
{{- range .Spec.Deployments}}
{{- range $key, $value := . }}
{{ $key }}:
  image:
    repository: {{ $value.Image }}
    tag: {{ $value.Tag }}
` + workloadTmpl +
	`
  ## {{ $key }} pods' resource requests and limits
  ## ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##  
  {{- if $value.Resources }}
  resources:
    # We usually recommend not to specify default resources and to leave this as a conscious
    # choice for the user. This also increases chances charts run on environments with little
    # resources, such as Minikube. If you do want to specify resources, uncomment the following
    # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
    limits: {}
    #   cpu: 100m
    #   memory: 128Mi
    requests: {}
    #   cpu: 100m
    #   memory: 128Mi
  {{- else }}
  resources: {}
  {{- end }}  # Resources
{{- end}}
{{- end}}
{{- end}}
{{- if .Spec.Jobs}}
{{- range .Spec.Jobs}}
{{- range $key, $value := . }}
{{ $key }}:
  image:
    repository: {{ $value.Image }}
    tag: {{ $value.Tag }}
` + workloadTmpl +
	`
  ## {{ $key }} pods' resource requests and limits
  ## ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##  
  {{- if $value.Resources }}
  resources:
    # We usually recommend not to specify default resources and to leave this as a conscious
    # choice for the user. This also increases chances charts run on environments with little
    # resources, such as Minikube. If you do want to specify resources, uncomment the following
    # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
    limits: {}
    #   cpu: 100m
    #   memory: 128Mi
    requests: {}
    #   cpu: 100m
    #   memory: 128Mi
  {{- else }}
  resources: {}
  {{- end }}  # Resources
{{- end}}
{{- end}}
{{- end}}
## Specify a imagePullPolicy
## Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent'
## ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images
##
pullPolicy: IfNotPresent
service:
  {{- if .Spec.Services}}
  {{- range .Spec.Services}}
  {{- range $key, $value := . }}
  ## {{ snakecase $key }} service definition
  {{ snakecase $key }}:
    type: ClusterIP
    port: 8080
  {{- end}}
  {{- end}}
  {{- end}}

## Container Security Context
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
##
containerSecurityContext:
  enabled: true
  runAsUser: 1001

livenessProbe:
  enabled: true
  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  successThreshold: 1
  failureThreshold: 5
readinessProbe:
  enabled: true
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 10
  successThreshold: 1
  failureThreshold: 5

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
`

// WriteValues outputs the values.yaml
func WriteValues(application models.Application, outDir string) {
	glog.Infof("Writing values.yaml")
	utils.OutputTemplate(application, valuesTmpl, filepath.Join(outDir, "values.yaml"))
}
