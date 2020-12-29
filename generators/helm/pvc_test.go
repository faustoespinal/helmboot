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

const pvcTestManifest = `
apiVersion: helmboot/beta/v1
type: application  # application, microservice, job
name: myapp
description: "This is a deployment of my awesome application"
version: 1.0.0
appVersion: 2.0.1
spec:
  storage:
  - mystorage1:
      size: 2Gi
      mode: ReadWriteOnce
  - mystorage2:
      size: 1Gi
      mode: ReadWriteMany
      storageClass: shared
`

func TestWritePvcs(t *testing.T) {
	t.Log("TestWritePvcs....")
	// Create the output directory
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	var application models.Application
	yamlFile := []byte(pvcTestManifest)
	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		glog.Errorf("Error parsing file: %v", err)
		panic(err)
	}

	metaApp := models.CreateMetaApplication(application)
	WritePvcs(metaApp, dir)

	// Verify file got created.
	filePath := path.Join(dir, "pvcs.yaml")
	t.Logf("Chart content: %s\n", filePath)
	assert.True(t, utils.FileExists(filePath))
}
