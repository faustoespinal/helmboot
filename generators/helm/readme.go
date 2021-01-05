package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

const readmeTmplHeader = `
# {{ .Name }}

## Overview
{{ .Description }}

` + "```bash" + `
# Testing configuration
$ helm repo add ehs https://cp-nexus-0.novalocal
$ helm install my-{{ .Name }} ehs/{{ .Name }}
` + "```" + `

` + "```bash" + `
# Production configuration
$ helm repo add ehs https://charts.bitnami.com/bitnami
$ helm install my-{{ .Name }} ehs/{{ .Name }} --values additional_values.yaml
` + "```" + `

## Introduction

This chart bootstraps a [{{ .Name }}](https://github.com/{{ .Name }}) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

Helm boot charts can be used with [Kubeapps](https://kubeapps.com/) for deployment and management of Helm Charts in clusters.


## Prerequisites

- Kubernetes 1.12+
- Helm 3.0+
- PV provisioner support in the underlying infrastructure

## Installing the Chart

To install the chart with the release name "my-release":

` + "```bash" + `
$ helm install my-{{ .Name }} ehs/{{ .Name }}
` + "```" + `

The command deploys {{ .Name }} on the Kubernetes cluster in the default configuration. The [Parameters](#parameters) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using "helm list"

## Uninstalling the Chart

To uninstall/delete the "my-{{ .Name }}" deployment:

` + "```bash" + `
$ helm delete my-{{ .Name }}
` + "```" + `

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the Redis chart and their default values.

`

const readmeTmplParams = `
| Parameter                                     | Description                                                                                                                                         | Default                                                 |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| global.imageRegistry                        | Global Docker image registry                                                                                                                        | "nil"                                                   |
| global.imagePullSecrets                     | Global Docker registry secret names as an array                                                                                                     | [] (does not add image pull secrets to deployed pods) |
| global.storageClass                         | Global storage class for dynamic provisioning                                                                                                       | "nil"                                                   |
| global.redis.password                       | Redis password (overrides password)                                                                                                               | "nil"                                                   |
`

const readmeTmplFooter = `

Specify each parameter using the "--set key=value[,key=value]" argument to "helm install". For example,

` + "```bash" + `
$ helm install my-{{ .Name }} --set {{ .Name }}.key="thevalue" ehs/{{ .Name }}
` + "```" + `

The above command sets the release {{ .Name }} helm chart value {{ .Name }}.key to "thevalue".

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

` + "```bash" + `
$ helm install my-{{ .Name }} -f values.yaml ehs/{{ .Name }}
` + "```"

var commonValues = []models.ChartValue{
	{Name: "pullPolicy",
		Description:  "Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent' (ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images)",
		DefaultValue: "IfNotPresent",
	},
	{Name: "registry",
		Description:  "Registry locator for the helm docker images",
		DefaultValue: "nul",
	},
	{Name: "containerSecurityContext.enabled",
		Description:  "Whether to enable settings enforcing container security context",
		DefaultValue: "true",
	},
	{Name: "containerSecurityContext.runAsUser",
		Description:  "The id of the user to use when running under a security context",
		DefaultValue: "1001",
	},
	{Name: "serviceAccount.create",
		Description:  "Whether to create a service account for this helm deployment",
		DefaultValue: "false",
	},
	{Name: "serviceAccount.name",
		Description:  "The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template",
		DefaultValue: "nul",
	},
	{Name: "livenessProbe.enabled",
		Description:  "Whether to enable liveness probe on container",
		DefaultValue: "true",
	},
	{Name: "livenessProbe.initialDelaySeconds",
		Description:  "Number of seconds after the container has started before liveness or readiness probes are initiated. Defaults to 0 seconds. Minimum value is 0.",
		DefaultValue: "30",
	},
	{Name: "livenessProbe.periodSeconds",
		Description:  "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
		DefaultValue: "10",
	},
	{Name: "livenessProbe.timeoutSeconds",
		Description:  "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.",
		DefaultValue: "5",
	},
	{Name: "livenessProbe.successThreshold",
		Description:  "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup Probes. Minimum value is 1.",
		DefaultValue: "1",
	},
	{Name: "livenessProbe.failureThreshold",
		Description:  "When a probe fails, Kubernetes will try failureThreshold times before giving up. Giving up in case of liveness probe means restarting the container. In case of readiness probe the Pod will be marked Unready. Defaults to 3. Minimum value is 1.",
		DefaultValue: "5",
	},
	{Name: "readinessProbe.enabled",
		Description:  "Whether to enable liveness probe on container",
		DefaultValue: "true",
	},
	{Name: "readinessProbe.initialDelaySeconds",
		Description:  "Number of seconds after the container has started before liveness or readiness probes are initiated. Defaults to 0 seconds. Minimum value is 0.",
		DefaultValue: "30",
	},
	{Name: "readinessProbe.periodSeconds",
		Description:  "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
		DefaultValue: "10",
	},
	{Name: "readinessProbe.timeoutSeconds",
		Description:  "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.",
		DefaultValue: "5",
	},
	{Name: "readinessProbe.successThreshold",
		Description:  "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup Probes. Minimum value is 1.",
		DefaultValue: "1",
	},
	{Name: "readinessProbe.failureThreshold",
		Description:  "When a probe fails, Kubernetes will try failureThreshold times before giving up. Giving up in case of liveness probe means restarting the container. In case of readiness probe the Pod will be marked Unready. Defaults to 3. Minimum value is 1.",
		DefaultValue: "5",
	},
}

func getValues(workloads []map[string]models.ContainerWorkload) []models.ChartValue {
	values := make([]models.ChartValue, 0)
	for _, workloadMap := range workloads {
		for workloadName := range workloadMap {
			values = append(values, models.ChartValue{
				Name:         workloadName + ".image.repository",
				Description:  "Container image repository",
				DefaultValue: "nul",
			})
			values = append(values, models.ChartValue{
				Name:         workloadName + ".image.tag",
				Description:  "Version tag of the container image for the " + workloadName + ".",
				DefaultValue: "nul",
			})
			values = append(values, models.ChartValue{
				Name:         workloadName + ".affinity",
				Description:  "Affinity for pod assignment (https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity",
				DefaultValue: "{}",
			})
			values = append(values, models.ChartValue{
				Name:         workloadName + ".nodeSelector",
				Description:  "Node labels for pod assignment (https://kubernetes.io/docs/user-guide/node-selection/)",
				DefaultValue: "{}",
			})
			values = append(values, models.ChartValue{
				Name:         workloadName + ".tolerations",
				Description:  "Tolerations for pod assignment (https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)",
				DefaultValue: "{}",
			})
			values = append(values, models.ChartValue{
				Name:         workloadName + ".resources",
				Description:  workloadName + " pods' resource requests and limits",
				DefaultValue: "{}",
			})
		}
	}
	return values
}

// WriteReadmeMd writes the Readme.md file for the HELM chart
func WriteReadmeMd(application models.Application, outDir string) {
	values := make([]models.ChartValue, 0)
	values = append(values, commonValues...)
	for _, svcMap := range application.Spec.Services {
		for svcName := range svcMap {
			values = append(values, models.ChartValue{
				Name:         "service." + svcName + ".type",
				Description:  "One of ClusterIP, NodePort or LoadBalancer",
				DefaultValue: "ClusterIP",
			})
			values = append(values, models.ChartValue{
				Name:         "service." + svcName + ".port",
				Description:  "Port number",
				DefaultValue: "8080",
			})
		}
	}
	values = append(values, getValues(application.Spec.Deployments)...)
	values = append(values, getValues(application.Spec.Jobs)...)
	params := readmeTmplParams
	for _, cvals := range values {
		params = params + "| " + cvals.Name + " | " + cvals.Description + " | " + cvals.DefaultValue + " |\n"
	}
	readmeHeader := utils.SOutputTemplate(application, readmeTmplHeader)
	readmeFooter := utils.SOutputTemplate(application, readmeTmplFooter)
	readme := readmeHeader + params + readmeFooter
	utils.WriteFile(readme, filepath.Join(outDir, "README.md"))
}
