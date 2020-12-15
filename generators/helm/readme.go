package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"
)

const readmeTmpl = `
# {{.Name}}

[Redis](http://redis.io/) is an advanced key-value cache and store. It is often referred to as a data structure server since keys can contain strings, hashes, lists, sets, sorted sets, bitmaps and hyperloglogs.

## Overview

` + "```bash" + `
# Testing configuration
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm install my-release bitnami/redis
` + "```" + `

` + "```bash" + `
# Production configuration
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm install my-release bitnami/redis --values values-production.yaml
` + "```" + `

## Introduction

This chart bootstraps a [Redis](https://github.com/bitnami/bitnami-docker-redis) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

Bitnami charts can be used with [Kubeapps](https://kubeapps.com/) for deployment and management of Helm Charts in clusters. This chart has been tested to work with NGINX Ingress, cert-manager, fluentd and Prometheus on top of the [BKPR](https://kubeprod.io/).

### Choose between Redis Helm Chart and Redis Cluster Helm Chart

You can choose any of the two Redis Helm charts for deploying a Redis cluster.
While [Redis Helm Chart](https://github.com/bitnami/charts/tree/master/bitnami/redis) will deploy a master-slave cluster using Redis Sentinel, the [Redis Cluster Helm Chart](https://github.com/bitnami/charts/tree/master/bitnami/redis-cluster) will deploy a Redis Cluster topology with sharding.
The main features of each chart are the following:

| Redis                                         | Redis Cluster                                                                                                                                       |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| Supports multiple databases                   | Supports only one database. Better if you have a big dataset                                                                                        |
| Single write point (single master)            | Multiple write points (multiple masters)                                                                                                            |
| ![Redis Topology](img/redis-topology.png)     | ![Redis Cluster Topology](img/redis-cluster-topology.png)                                                                                           |

## Prerequisites

- Kubernetes 1.12+
- Helm 3.0-beta3+
- PV provisioner support in the underlying infrastructure

## Installing the Chart

To install the chart with the release name "my-release":

` + "```bash" + `
$ helm install my-release bitnami/redis
` + "```" + `

The command deploys Redis on the Kubernetes cluster in the default configuration. The [Parameters](#parameters) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using "helm list"

## Uninstalling the Chart

To uninstall/delete the "my-release" deployment:

` + "```bash" + `
$ helm delete my-release
` + "```" + `

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the Redis chart and their default values.

| Parameter                                     | Description                                                                                                                                         | Default                                                 |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| global.imageRegistry                        | Global Docker image registry                                                                                                                        | "nil"                                                   |
| global.imagePullSecrets                     | Global Docker registry secret names as an array                                                                                                     | [] (does not add image pull secrets to deployed pods) |
| global.storageClass                         | Global storage class for dynamic provisioning                                                                                                       | "nil"                                                   |
| global.redis.password                       | Redis password (overrides password)                                                                                                               | "nil"                                                   |
| image.registry                              | Redis Image registry                                                                                                                                | docker.io                                             |
| image.repository                            | Redis Image name                                                                                                                                    | bitnami/redis                                         |
| image.tag                                   | Redis Image tag                                                                                                                                     | {TAG_NAME}                                            |
| image.pullPolicy                            | Image pull policy                                                                                                                                   | IfNotPresent                                          |
| image.pullSecrets                           | Specify docker-registry secret names as an array                                                                                                    | "nil"                                                   |
| nameOverride                                | String to partially override redis.fullname template with a string (will prepend the release name)                                                  | "nil"                                                   |
| fullnameOverride                            | String to fully override redis.fullname template with a string                                                                                      | "nil"                                                   |
| cluster.enabled                             | Use master-slave topology                                                                                                                           | true                                                  |
| cluster.slaveCount                          | Number of slaves                                                                                                                                    | 2                                                     |
| existingSecret                              | Name of existing secret object (for password authentication)                                                                                        | "nil"                                                   |
| existingSecretPasswordKey                   | Name of key containing password to be retrieved from the existing secret                                                                            | "nil"                                                   |
| usePassword                                 | Use password                                                                                                                                        | true                                                  |
| usePasswordFile                             | Mount passwords as files instead of environment variables                                                                                           | false                                                 |
| password                                    | Redis password (ignored if existingSecret set)                                                                                                      | Randomly generated                                      |
| configmap                                   | Additional common Redis node configuration (this value is evaluated as a template)                                                                  | See values.yaml                                         |
| clusterDomain                               | Kubernetes DNS Domain name to use                                                                                                                   | cluster.local                                         |
| networkPolicy.enabled                       | Enable NetworkPolicy                                                                                                                                | false                                                 |
| podSecurityPolicy.create                    | Specifies whether a PodSecurityPolicy should be created                                                                                             | false                                                 |

Specify each parameter using the "--set key=value[,key=value]" argument to "helm install". For example,

` + "```bash" + `
$ helm install my-release \
  --set password=secretpassword \
    bitnami/redis
` + "```" + `

The above command sets the Redis server password to "secretpassword".

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

` + "```bash" + `
$ helm install my-release -f values.yaml bitnami/redis
` + "```"

// WriteReadmeMd writes the Readme.md file for the HELM chart
func WriteReadmeMd(application models.Application, outDir string) {
	utils.OutputTemplate(application, readmeTmpl, filepath.Join(outDir, "README.md"))

}
