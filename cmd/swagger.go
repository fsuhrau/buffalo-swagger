// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"

	"github.com/fsuhrau/buffalo-swagger/generator"
	"github.com/fsuhrau/buffalo-swagger/parser"
	"github.com/spf13/cobra"
)

var yamlExport bool

// swaggerCmd represents the swagger command
var swaggerCmd = &cobra.Command{
	Use:     "swagger",
	Aliases: []string{"s"},
	Short:   "Tool to generate a swagger file.",
	Run: func(cmd *cobra.Command, args []string) {
		path := ""
		if len(args) > 0 {
			path = args[0]
		}
		outputFile := ""
		if len(args) > 1 {
			outputFile = args[1]
		}
		parser := parser.NewParser(path)
		err := parser.ParseProject()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		gen := generator.NewGenerator(outputFile)
		err = gen.Generate(parser, yamlExport)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(swaggerCmd)
	swaggerCmd.Flags().BoolVarP(&yamlExport, "yaml", "y", false, "export as yaml")
}
