# helmboot
[![TravisCI](https://travis-ci.com/faustoespinal/helmboot.svg?branch=main)](https://travis-ci.com/github/faustoespinal/helmboot)
[![Go Report Card](https://goreportcard.com/badge/github.com/faustoespinal/helmboot)](https://goreportcard.com/report/github.com/faustoespinal/helmboot)

helmboot can be used as a standalone command or as a helm plugin to quickly scaffold helm charts for kubernetes cloud-native applications.

## Basic Usage

The following command generates the helm chart based on the helmboot app descriptor file cn-application.yaml.
```
helmboot create --workload sample-apps/cn-application.yaml --output charts/
```

## Examples

- [bookinfo app](https://github.com/faustoespinal/helmboot/blob/main/sample-apps/bookinfo-app.yaml)
- [emojivoto app](https://github.com/faustoespinal/helmboot/blob/main/sample-apps/emojivoto-app.yaml)

