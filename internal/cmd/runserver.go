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
	"golculator/internal/server"

	"github.com/spf13/cobra"
)

// runserverCmd represents the runserver command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Run a client with given port",
	Long: `Run a client with given port, and listen to the port.
	for example: golculator runserver 8080
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runserver called")
		server := server.SocketServer{
			Port: args[0],
		}
		server.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
