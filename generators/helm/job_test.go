package helm

import (
	"helmboot/models"
	"helmboot/utils"
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

const jobTestManifest = `
apiVersion: helmboot/beta/v1
type: application  # application, microservice, job
name: myapp
description: "This is a deployment of my awesome application"
version: 1.0.0
appVersion: 2.0.1
spec:
  jobs:
    - taskjob:
        image: somejobimage
        tag: 1.0.1
        command: "echo Initializing; ls -ls"
        configmaps:
          - appconfig
        env:
          - name: INIT_THE_MESSAGE
            value: "Hello there again"
        databases:
          - my-db-connection
        storage:
          - mystorage:
              mount: "/mnt/storage"
  storage:
    - mystorage:
        size: 1Gi
        mode: ReadWriteMany
        storageClass: shared
  configmaps:
    - appconfig:
        data:
          - mykey: myvalue
  databases:
    - my-db-connection
`

func TestWriteJobs(t *testing.T) {
	t.Log("TestWriteJobs....")
	// Create the output directory
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	var application models.Application
	yamlFile := []byte(jobTestManifest)
	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		zap.S().Errorf("Error parsing file: %v", err)
		panic(err)
	}

	metaApp := models.CreateMetaApplication(application)
	WriteJobs(metaApp, dir)

	// Verify file got created.
	filePath := path.Join(dir, "jobs.yaml")
	t.Logf("Chart content: %s\n", filePath)
	assert.True(t, utils.FileExists(filePath))
}
