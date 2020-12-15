package generators

import "helmboot/models"

// Templator interface abstracting a module that can issue ....
type Templator interface {
	// Descriptive name of the template generator
	Name() string

	// Write the templates to the specified output directory
	Write(application models.Application, outDir string)
}
