package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func touchFile(t *testing.T, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
}

func TestCreateDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestOutputDir: %v\n", dir)
	directory := path.Join(dir, "outpath")
	CreateDir(directory)

	_, err = os.Stat(directory)
	dirExists := !os.IsNotExist(err)
	assert.True(t, dirExists)
	t.Logf("Directory Exists: %v", dirExists)
}

func TestClearDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		t.Fatal(err)
	}

	touchFile(t, path.Join(dir, "temp1.txt"))
	touchFile(t, path.Join(dir, "temp2.txt"))
	touchFile(t, path.Join(dir, "temp3.txt"))

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("There should be 3 files: %v %v\n", dir, len(files))
	assert.True(t, len(files) == 3)

	// delete the directory
	ClearDir(dir)

	// Verify directory is empty
	files, err = ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("Content after delete: %v\n", len(files))
	assert.True(t, len(files) == 0)
}

type TestType struct {
	Hello string
	There string
}

func TestSOutputTemplate(t *testing.T) {
	value := TestType{
		Hello: "hello",
		There: "there",
	}
	templateString := "{{ .Hello }}-{{.There}}"
	renderedTemplate := SOutputTemplate(value, templateString)
	t.Logf("SRenderedTemplate: %v", renderedTemplate)
	assert.True(t, renderedTemplate == "hello-there")
}

func TestOutputTemplate(t *testing.T) {
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		t.Fatal(err)
	}
	filePath := path.Join(dir, "temp.txt")

	value := TestType{
		Hello: "hello",
		There: "there",
	}
	templateString := "{{ .Hello }}-{{.There}}"
	OutputTemplate(value, templateString, filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal("err")
	}
	renderedTemplate := string(content)
	t.Logf("RenderedTemplate: %v", renderedTemplate)
	assert.True(t, renderedTemplate == "hello-there")
}
