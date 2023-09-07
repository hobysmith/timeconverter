// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package cmd

import (
	"fmt"
	"github.com/hobysmith/timeconverter/helpers"
	"github.com/spf13/cobra"
)

var clearLocal bool
var clearGlobal bool

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Removes local and global config settings",
	Long:  "Removes local and global config settings",
	Run: func(cmd *cobra.Command, args []string) {
		if !clearLocal && !clearGlobal {
			_ = cmd.Help()
			return
		}

		if clearLocal {
			err := helpers.CmdHelpers.ClearLocalConfig()
			if err != nil {
				fmt.Printf("Error clearing local config: %s\n", err)
				return
			}
			fmt.Println("Local config removed")
		}

		if clearGlobal {
			err := helpers.CmdHelpers.ClearGlobalConfig()
			if err != nil {
				fmt.Printf("Error clearing global config: %s\n", err)
				return
			}
			fmt.Println("Global config removed")
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
	clearCmd.Flags().BoolVarP(&clearLocal, "local", "l", false, "If true, will remove the local config file.  Default is false.")
	clearCmd.Flags().BoolVarP(&clearGlobal, "global", "g", false, "If true, will remove the global config file.  Default is false.")
}
