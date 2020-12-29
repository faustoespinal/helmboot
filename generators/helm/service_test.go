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

const svcTestManifest = `
apiVersion: helmboot/beta/v1
type: application
name: myapp
description: "This is a deployment services"
version: 1.0.0
appVersion: 2.0.1
spec:
  services:
  - svc1:
      deployment: deadbeef
  - svc2:
      deployment: crazycow
`

func TestWriteServices(t *testing.T) {
	t.Log("TestWriteServices....")
	// Create the output directory
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	var application models.Application
	yamlFile := []byte(svcTestManifest)
	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		glog.Errorf("Error parsing file: %v", err)
		panic(err)
	}

	metaApp := models.CreateMetaApplication(application)
	WriteServices(metaApp, dir)

	// Verify file got created.
	filePath := path.Join(dir, "service.yaml")
	t.Logf("Chart content: %s\n", filePath)
	assert.True(t, utils.FileExists(filePath))
}
