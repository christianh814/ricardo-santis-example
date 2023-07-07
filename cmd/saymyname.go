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

	"github.com/manifoldco/promptui"
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

		// Watch for changes in the config file and write them if prompt changed the value
		if ask {
			prompt := promptui.Prompt{
				Label: "Your name",
			}

			result, err := prompt.Run()

			if err != nil {
				log.Fatalf("Prompt failed %v\n", err)
			}
			viper.WatchConfig()
			viper.Set("name", result)
			viper.WriteConfig()
		}

		// Grab the name var
		name := viper.GetString("name")

		// If name is empty, exit with error
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

	// Look for "name" in other places. 1st from CLI, 2nd Env Var, 3rd config file
	viper.BindPFlag("name", saymynameCmd.Flags().Lookup("name"))
}
