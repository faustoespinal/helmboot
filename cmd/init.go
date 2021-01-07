/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"helmboot/appmodels"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an application",
	Long:  `Initialize an application descriptor from a standard template.`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.Info("Initialize an application descriptor.")

		isWebService, _ := cmd.Flags().GetBool("webservice")
		isTask, _ := cmd.Flags().GetBool("task")
		appName, err := cmd.Flags().GetString("name")
		if err != nil {
			cmd.Help()
			os.Exit(1)
		}

		glog.Infof("** Init: %v %v %v\n", isWebService, isTask, appName)
		appType := "generic"
		if isWebService {
			appType = "web"
		} else if isTask {
			appType = "task"
		}
		appmodels.GenerateAppDescriptor(appName, appType)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.PersistentFlags().Bool("webservice", false, "Use the web service template.")
	initCmd.PersistentFlags().Bool("task", false, "Use the task application template.")
	initCmd.PersistentFlags().String("name", "", "Name of the application to initialize.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
