package cmd

import (
	"helmboot/utils"
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const simpleApp = `
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

func TestCreate(t *testing.T) {
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("OutputDir: %v\n", dir)

	performCreate([]byte(simpleApp), dir, true)
	chartFilePath := path.Join(dir, "simple-app", "Chart.yaml")
	t.Logf("Chart content: %s\n", chartFilePath)
	assert.True(t, utils.FileExists(chartFilePath))
}
