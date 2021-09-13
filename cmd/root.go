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
	"fmt"

	"github.com/RaymondArias/SecretSource/pkg/secretreader"
	"github.com/spf13/cobra"
)

var sourceFile string
var secretStore string
var region string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "env",
	Short: "Reads secret values from file and replaces them with secret values",
	Long: `Reads secret values from file and replaces them with secret value
	Useful for populating environment variables`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if sourceFile == "" {
			fmt.Printf("No file passed in\n")
			return
		}
		var secStore secretreader.SecretStore
		switch secretStore {
		case "aws":
			secretStore, err := secretreader.NewSSMReader(region)
			if err != nil {
				fmt.Printf("Error creating SSM Reader: %s\n", err.Error())
				return
			}
			secStore = secretStore
		default:
			fmt.Println("No secret storage passed in")
		}
		secReader := secretreader.NewSecretReader(secStore)
		err := secReader.GenerateSource(sourceFile)
		if err != nil {
			fmt.Printf("error generating env variables: %s", err.Error())
			fmt.Println()
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&sourceFile, "file", "f", "", "source file with secrets")
	rootCmd.Flags().StringVarP(&secretStore, "secret-store", "s", "aws", "secret storage implementation, only AWS at the momemt")
	rootCmd.Flags().StringVarP(&region, "region", "r", "us-west-2", "aws region were parameters are stored")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
