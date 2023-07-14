/*
Copyright © 2023 Christian Hernandez <christian@chernand.io>

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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// saymynameCmd represents the saymyname command
var saymynameCmd = &cobra.Command{
	Use:   "saymyname",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Grab the ask variable
		ask, _ := cmd.Flags().GetBool("ask")

		if ask {
			// Prompt the user for their name if they provided the --ask flag
			var result string
			fmt.Printf("Your name: ")

			// Scan and check for errors
			_, err := fmt.Scanln(&result)
			if err != nil {
				log.Fatal(err)
			}

			// Report error if user entered nothing
			if result == "" {
				log.Fatal("Name cannot be blank")
			}

			// Watch for changes in the config
			viper.WatchConfig()

			// Set the name var in the config file based on the result from the prompt
			viper.Set("name", result)

			// Write the config file with the new value
			viper.WriteConfig()
		}

		// Grab the name var from any of the 3 places (CLI, Env Var, Config File)
		name := viper.GetString("name")

		// If name is empty, exit with error. Else, print the name
		if name == "" {
			log.Fatal("Name is required")
		} else {
			fmt.Println(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(saymynameCmd)
	saymynameCmd.Flags().StringP("name", "n", "", "Your name")
	saymynameCmd.Flags().BoolP("ask", "a", false, "Ask for your name")

	// Look for "name" in other places. In order or precedence: 1st from CLI, 2nd Env Var, 3rd config file
	viper.BindPFlag("name", saymynameCmd.Flags().Lookup("name"))
}
