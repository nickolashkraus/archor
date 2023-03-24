// Copyright 2023 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Mirror an existing RSS feed and content",
	Long: `Mirroring simply mirrors an existing RSS feed.

It creates a one-to-one duplication of the upstream RSS feed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mirror called")
	},
}

func init() {
	archorCmd.AddCommand(mirrorCmd)
	mirrorCmdHelp := `Destination of the RSS feed and content.
Currently, this can be a directory on the filesystem
(ex. path/to/dir) or an S3 bucket URI (ex. s3://my-bucket).
Defaults to the current working directory.`
	mirrorCmd.Flags().StringP("destination", "d", ".", mirrorCmdHelp)
}
