// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/fanux/fist/terminal"
	"github.com/spf13/cobra"
)

// terminalCmd represents the terminal command
var terminalCmd = &cobra.Command{
	Use:   "terminal",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		terminal.Serve(cmd)
	},
}

func init() {
	rootCmd.AddCommand(terminalCmd)

	// Here you will define your flags and configuration settings.
	terminalCmd.Flags().Uint16P("port", "P", 8080, "start  listening port")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// terminalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// terminalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
