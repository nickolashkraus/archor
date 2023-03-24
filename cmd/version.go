// Copyright 2023 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Archor",
	Long:  `All software has versions. This is Archor's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Archor v0.0.1")
	},
}

func init() {
	archorCmd.AddCommand(versionCmd)
}
