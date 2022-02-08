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

const ingressTestManifest = `
apiVersion: helmboot/beta/v1
type: application
name: myapp
description: "This is a deployment of ingresses for an app"
version: 1.0.0
appVersion: 2.0.1
spec:
  services:
  - svc1:
      deployment: deadbeef
  - svc2:
      deployment: crazycow
  ingresses:
  - svc1-ingress:
      service: svc1
  - svc2-ingress:
      service: svc2
  - svc3-ingress:
      service: svc3
      namespace: dicom
      externalService: dcm-stow
`

func TestWriteIngresses(t *testing.T) {
	t.Log("TestWriteIngresses....")
	// Create the output directory
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	var application models.Application
	yamlFile := []byte(ingressTestManifest)
	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		zap.S().Errorf("Error parsing file: %v", err)
		panic(err)
	}

	metaApp := models.CreateMetaApplication(application)
	WriteIngresses(metaApp, dir)

	// Verify file got created.
	filePath := path.Join(dir, "ingress.yaml")
	t.Logf("Chart content: %s\n", filePath)
	assert.True(t, utils.FileExists(filePath))
}
