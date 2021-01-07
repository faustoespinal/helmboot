package appmodels

import (
	"fmt"
	"helmboot/utils"
)

const webApp = `
    webapp {{ .Name }}
`

const taskApp = `
    task {{ .Name }}
`

const genericApp = `
    generic {{ .Name }}
`

// AppInfo contains the metadata of an application
type AppInfo struct {
	Name string
}

// GenerateAppDescriptor prints an application template
func GenerateAppDescriptor(name string, apptype string) {
	ai := AppInfo{
		Name: name,
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
