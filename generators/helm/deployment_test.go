package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/golang/glog"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const deploymentTestManifest = `
apiVersion: helmboot/beta/v1
type: application  # application, microservice, job
name: myapp
description: "This is a deployment of my awesome application"
version: 1.0.0
appVersion: 2.0.1
spec:
  deployments:
  - deadbeef:
      image: someimage
      tag: 1.0.0
      configmaps:
        - appconfig
      env:
        - name: INIT_MESSAGE
          value: "Hello there"
      ports: 
        - containerPort: 8080
          name: deadbeef-http
        - containerPort: 8081
      databases:
        - my-db-connection
      resources:
        requests:
            memory: "64Mi"
            cpu: "250m"
        limits:
            memory: "128Mi"
            cpu: "500m"
  configmaps:
    - appconfig:
        data:
          - mykey: myvalue
  databases:
    - my-db-connection
`

func TestWriteDeployments(t *testing.T) {
	t.Log("TestWriteDeployments....")
	// Create the output directory
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	var application models.Application
	yamlFile := []byte(deploymentTestManifest)
	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		glog.Errorf("Error parsing file: %v", err)
		panic(err)
	}

	metaApp := models.CreateMetaApplication(application)
	WriteDeployments(metaApp, dir)

	// Verify file got created.
	filePath := path.Join(dir, "deployment.yaml")
	t.Logf("Chart content: %s\n", filePath)
	assert.True(t, utils.FileExists(filePath))
}
