package appmodels

import (
	"fmt"
	"helmboot/utils"
)

const webApp = `
apiVersion: helmboot/beta/v1
type: application
name: {{ .Name }}
description: "Sample web service application"
version: {{ .Version }}
appVersion: {{ .AppVersion }}
spec:
  security:
    grantTypes: 
      - implicit
    roles:
      - admin:
        scopes:
          - scope1
          - scope2
  testing:
    image: curlimages/curl:7.74.0
    command: ['curl']
  deployments:
    - {{ .Name }}-deployment:
      image: someimage
      tag: 1.0.0
      configmaps:
        - app-config
      secrets:
        - app-secret
      env:
        - name: SAMPLE_ENV
          value: "ExampleValue"
      ports: 
        - containerPort: 8080
          name: deadbeef-http
      databases:
        - db-connection
      resources:
        requests:
          memory: "64Mi"
          cpu: "250m"
        limits:
          memory: "128Mi"
          cpu: "500m"
      messaging:
        - msg-queue
      storage:
        - storage1:
          mount: "/mnt/store1"
        - storage2:
          mount: "/mnt/store2"
  services:
    - {{ .Name }}-svc:
      deployment: {{ .Name }}-deployment
  storage:
    - storage1:
      size: 2Gi
      mode: ReadWriteOnce
    - storage2:
      size: 1Gi
      mode: ReadWriteMany
      storageClass: shared
  ingresses:
    - {{ .Name }}-ingress:
      service: {{ .Name }}-svc
    - {{ .Name }}-ext-ingress:
      service: {{ .Name }}-ext-svc
      namespace: edison-core
      externalService: eis-stow
  configmaps:
    - app-config:
      data:
        - mykey: myvalue
  secrets:
    - app-secret:
      type: "opaque"
      data:
        - SECRET_A
        - SECRET_B
  databases:
    - db-connection
  messaging:
    - msg-queue	
`

const taskApp = `
apiVersion: helmboot/beta/v1
type: application
name: {{ .Name }}
description: "Sample web service application"
version: {{ .Version }}
appVersion: {{ .AppVersion }}
spec:
  security:
    grantTypes: 
      - implicit
    roles:
      - admin:
        scopes:
          - scope1
          - scope2
  jobs:
    - {{ .Name }}-job:
        image: task-image
        tag: 1.0.1
        command: "echo Initializing; ls -ls"
        configmaps:
          - app-config
        secrets:
          - app-secret
        env:
          - name: INIT_THE_MESSAGE
            value: "Hello there again"
        databases:
          - db-connection
        messaging:
          - msg-queue
        storage:
          - {{ .Name }}-storage:
              mount: "/mnt/storage"      
  storage:
    - {{ .Name }}-storage:
      size: 2Gi
      mode: ReadWriteOnce
  configmaps:
    - app-config:
      data:
        - mykey: myvalue
  secrets:
    - app-secret:
      type: "opaque"
      data:
        - SECRET_A
        - SECRET_B
  databases:
    - db-connection
  messaging:
    - msg-queue	
`

const genericApp = `
    generic {{ .Name }}
`

// AppInfo contains the metadata of an application
type AppInfo struct {
	Name       string
	Version    string
	AppVersion string
}

// GenerateAppDescriptor prints an application template
func GenerateAppDescriptor(name string, apptype string) {
	ai := AppInfo{
		Name:       name,
		Version:    "1.0.0",
		AppVersion: "1.0.0",
	}
	app := ""
	switch apptype {
	case "web":
		app = utils.SOutputTemplate(ai, webApp)
	case "task":
		app = utils.SOutputTemplate(ai, taskApp)
	case "generic":
		app = utils.SOutputTemplate(ai, genericApp)
	}
	fmt.Println(app)
}
