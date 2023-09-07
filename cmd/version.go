// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package cmd

import (
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints the app's version info",
	Long:    `Prints the app's version info`,
	Aliases: []string{},
	Run: func(cmd *cobra.Command, args []string) {
		printVersionInfo(false)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
