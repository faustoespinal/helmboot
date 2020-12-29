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

const secretTestManifest = `
apiVersion: helmboot/beta/v1
type: application  # application, microservice, job
name: myapp
description: "This is a deployment of my awesome application"
version: 1.0.0
appVersion: 2.0.1
spec:
  secrets:
  - appsecret1:
      type: "opaque"
      data:
        - SECRET_A
        - SECRET_B
  - appsecret2:
      data:
        - SECRET_C
        - SECRET_D
`

func TestWriteSecrets(t *testing.T) {
	t.Log("TestWriteSecrets....")
	// Create the output directory
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	var application models.Application
	yamlFile := []byte(secretTestManifest)
	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		glog.Errorf("Error parsing file: %v", err)
		panic(err)
	}

	metaApp := models.CreateMetaApplication(application)
	WriteSecrets(metaApp, dir)

	// Verify file got created.
	filePath := path.Join(dir, "secrets.yaml")
	t.Logf("Chart content: %s\n", filePath)
	assert.True(t, utils.FileExists(filePath))
}
