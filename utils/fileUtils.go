package utils

import (
	"html/template"
	"os"
	"path/filepath"
)

// CreateDir creates specified directory and return true if successful, false otherwise
func CreateDir(dir string) bool {
	// Create outDir if necessary
	stat, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0755)
		if err == nil {
			return true
		}
	} else {
		if stat.IsDir() {
			return true
		}
	}
	return false
}

// ClearDir removes all content from the specified directory
func ClearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteFile outputs the content string to the given file path
func WriteFile(content string, filePath string) error {
	outFile, err := os.Create(filePath)
	defer outFile.Close()
	if err != nil {
		return err
	}
	_, err = outFile.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// OutputTemplate evaluates an output template
func OutputTemplate(templateValues interface{}, templateContent string, filePath string) {
	outFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	tmpl, err := template.New("helm").Parse(templateContent)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(outFile, templateValues)
	if err != nil {
		panic(err)
	}
}

// FileExists returns true if a file exists and false otherwise
func FileExists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
