testpackage helm

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

const valuesTestManifest = `
apiVersion: helmboot/beta/v1
type: application  # application, microservice, job
name: simple-app
description: "This is a deployment of a simple application"
version: 1.0.0
appVersion: 2.0.1
spec:
  security:
    grantTypes:
      - implicit
    roles:
      - admin:
          scopes:
            - scope1
  deployments:
    - simplebeef:
        image: someimage
        tag: 1.0.0
        configmaps:
          - appconfig
        port: 8080
        databases:
          - my-db-connection
        storage:
          - mystorage1:
              mount: "/mnt/store1"
  services:
    - svc1:
        deployment: simplebeef
  storage:
    - mystorage1:
        size: 2Gi
        mode: ReadWriteOnce
  configmaps:
    - appconfig:
        data:
          - mykey: myvalue
`

func TestWriteValues(t *testing.T) {
	t.Log("TestWriteValues....")
	// Create the output directory
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	var application models.Application
	yamlFile := []byte(valuesTestManifest)
	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		glog.Errorf("Error parsing file: %v", err)
		panic(err)
	}

	WriteValues(application, dir)

	// Verify file got created.
	filePath := path.Join(dir, "values.yaml")
	t.Logf("Chart content: %s\n", filePath)
	assert.True(t, utils.FileExists(filePath))
}
