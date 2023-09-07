// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

// Package cmd contains all the Cobra CLI wrappers for the command line parsing and execution.
package cmd

import (
	"fmt"
	"github.com/hobysmith/timeconverter/converter"
	"github.com/hobysmith/timeconverter/helpers"
	"github.com/spf13/cobra"
	"os"
)

// This root file implements handling for the majority of the command flags.
// This is a typical Cobra CLI pattern.

// errInInit passes errors that can occur during the init function to the
// command functionality, where it can be handled more gracefully for the runtime.
var errInInit error

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tc",
	Short: "A utility for converting time values",
	Long:  `A utility for converting time values`,
	Args: func(cmd *cobra.Command, args []string) error {
		if helpers.CheckIsPiped() {
			return nil
		}

		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		// The first param should always be the input time
		helpers.CmdHelpers.Value = args[0]
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if errInInit != nil {
			if helpers.ExitCode == helpers.ExitCodeSuccess {
				// ExitCode was not set in LoadOutputPrinter(), so use general exit code here
				helpers.ExitCode = helpers.ExitCodeUnknownErrorInRootCommand
			}

			return errInInit
		}

		defer func() {
			outputErr := helpers.OP.UnloadOutputPrinter()
			if outputErr != nil {
				// we don't want to overwrite a flow through error state from the main app body,
				// so we only set return err if it's not set at this point
				if err == nil {
					// don't print anything here, because the caller will print an error message for us
					err = outputErr
				} else {
					// since we are not passing the error back, we'll print
					// a console out message here
					fmt.Printf("Error in UnloadOutputPrinter(): %s\n", err)
				}
			}
		}()

		err = converter.New().Convert(false)
		if err != nil {
			if helpers.ExitCode == helpers.ExitCodeSuccess {
				// ExitCode was not set in Convert(), so use general exit code here
				helpers.ExitCode = helpers.ExitCodeUnknownErrorInRootCommand
			}

			if !helpers.CmdHelpers.OutputValueOnly {
				fmt.Println(err)
			}

			helpers.CmdHelpers.ErrResult = err
			return nil
		}

		if helpers.CmdHelpers.SetDefault {
			helpers.CmdHelpers.SaveLocalDefault()
		}

		if helpers.CmdHelpers.SetGlobalDefault {
			helpers.CmdHelpers.SaveGlobalDefault()
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer func() {
		if r := recover(); r != nil {
			helpers.ExitCode = helpers.ExitCodePanicInExecute

			// Todo: do we want to honor the no output concept for piping as well relating to critical errors?
			if !helpers.CmdHelpers.OutputValueOnly {
				fmt.Printf("Panic recovered in cmd.Execute(): %s\n", r)
			}
		}
	}()

	if len(os.Args) == 1 && !helpers.CheckIsPiped() {
		_ = rootCmd.Help()
		return
	}

	err := rootCmd.Execute()
	if err != nil {
		helpers.ExitCode = helpers.ExitCodeErrorReturnedToExecute
	}
}

func init() {
	cmd := GetRootCmd()
	cmd.Use = "timeconverter dateTimeValue [flags]"
	cmd.Example = `  timeconverter 681678000 --input-format UnixSecs
  timeconverter now --output-format USTimeStampZ
  timeconverter 681678000 --input-format UnixMilli --output-format RFC3339
  timeconverter 681678000000 --input-format UnixMilli --input-format secs --output-format RFC3339
  timeconverter 681678000000 --input-format UnixMilli --output-format custom --output-layout "mmm yyyy-mm-dd hhh:nn:ss.000 zthhmm""
  timeconverter 681678000000 --input-format uNIxmilLI --output-format customGo --output-layout "Jan 2006-01-02 15:04:05.000 Z-0700"
  timeconverter show --time-formats
  timeconverter show --custom-entities`

	cmd.Flags().StringVarP(&helpers.CmdHelpers.OutputTargetName, "output-target", "t", "", "Indicates the type of output. Either console or clipboard.  If omitted, the default is console.")
	cmd.Flags().StringVarP(&helpers.CmdHelpers.InputFormatName, "input-format", "i", "USDateTimeZ", "The type of input. Use \"timeconverter show -time-formats\" for a list of allowed time formats.  Default is USDateTimeZ.")
	cmd.Flags().StringVarP(&helpers.CmdHelpers.InputLayout, "input-layout", "l", "", "When input format is set to \"custom\" or \"customgo\", this is the layout text.")
	cmd.Flags().StringVarP(&helpers.CmdHelpers.OutputFormatName, "output-format", "o", "USDateTimeZ", "The desired output format.  Use \"timeconverter show -time-formats\" for a list of allowed formats.  Default is USDateTimeZ.")
	cmd.Flags().StringVarP(&helpers.CmdHelpers.OutputLayout, "output-layout", "r", "", "When output format is \"custom\" or \"customgo\", this is the layout text.")
	cmd.Flags().BoolVarP(&helpers.CmdHelpers.OutputValueOnly, "output-value-only", "v", false, "If true, only the value will be sent to the output, with the possible exception of critical errors.  Default is false.")
	cmd.Flags().BoolVarP(&helpers.CmdHelpers.PipeMode, "piped", "p", false, "timeconverter detects piped input automatically, but IF that is not working for some reason, then this flag explicitly indicates that you are piping input in from another app. When true, input is read from stdin. By default, only the decoded value is printed to output, to support piping the timeconverter output to another app. That can be changed via flags.")
	cmd.Flags().BoolVarP(&helpers.CmdHelpers.SetGlobalDefault, "set-global-default", "", false, "When provided, a global config will be created, or updated if it already exists.")
	cmd.Flags().BoolVarP(&helpers.CmdHelpers.SetDefault, "set-default", "", false, "When provided, a local default config will be created.  The local config will override global configs when both are set.")
	cmd.Flags().StringVarP(&helpers.CmdHelpers.OutputTimeZone, "output-timezone", "z", "", "A timezone to use when converting the output time.  If not specified, the local time will be used for the output time. Can be a country/city ref from IANA TZ database or a timezone offset like -0700, +0000 or +0300")

	errInInit = helpers.LoadOutputPrinter()
	if errInInit != nil {
		return
	}
	helpers.CmdHelpers.LoadConfigIfExists()
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}
