/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"helmboot/generators/helm"
	"helmboot/models"
	"helmboot/utils"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func performCreate(yamlFile []byte, outDir string) {
	utils.CreateDir(outDir)

	var application models.Application
	//var job models.Job
	err := yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		glog.Errorf("Error parsing file: %v", err)
		panic(err)
	}

	workloadDir := filepath.Join(outDir, application.Name)
	utils.CreateDir(workloadDir)
	utils.ClearDir(workloadDir)

	generator := new(helm.Generator)
	generator.Write(application, workloadDir)
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.Info("Create called")

		outDir, err := cmd.Flags().GetString("output")
		if err != nil || len(outDir) <= 0 {
			cmd.Help()
			os.Exit(1)
		}

		// Get the input workload name..
		fileName, err := cmd.Flags().GetString("workload")
		if err != nil || len(fileName) <= 0 {
			cmd.Help()
			os.Exit(1)
		}

		fileName, _ = filepath.Abs(fileName)
		yamlFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		performCreate(yamlFile, outDir)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")
	createCmd.PersistentFlags().String("workload", "", "File name for the input application template")
	createCmd.PersistentFlags().String("output", "", "Directory name where to create the workload deployment (e.g. Helm)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
