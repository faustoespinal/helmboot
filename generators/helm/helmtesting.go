package helm

import (
	"fmt"
	"helmboot/models"
	"helmboot/utils"
	"path/filepath"

	"go.uber.org/zap"
)

const testingStub = `
apiVersion: v1
kind: Pod
metadata:
  name: test-{{ .TargetService }}
  labels:
    app.kubernetes.io/name: test-{{ .TargetService }}
spec:
  containers:
    - name: helm-test
      image: {{ .Application.Spec.Testing.Image }}
      command: {{ .Application.Spec.Testing.Command }}
      args:  [ "http://{{ .TargetService }}:80" ]
  restartPolicy: Never
`

// TestInfo input to test templates
type TestInfo struct {
	TargetService string             `yaml: "targetService"`
	Application   models.Application `yaml: "application"`
}

// WriteSvcTests outputs a test for a given service
func WriteSvcTests(metaApp models.MetaApplication, outDir string) {
	zap.S().Infof("Writing helm-svc testing")

	for _, svc := range metaApp.Application.Spec.Services {
		for svcKey := range svc {
			tinfo := TestInfo{
				TargetService: svcKey,
				Application:   metaApp.Application,
			}
			fileName := fmt.Sprintf("test-%s.yaml", svcKey)
			utils.OutputTemplate(tinfo, testingStub, filepath.Join(outDir, fileName))
		}
	}
}
