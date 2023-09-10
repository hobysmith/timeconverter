// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package cmd

import (
	"fmt"
	"github.com/hobysmith/timeconverter/helpers"
	"github.com/spf13/cobra"
)

var showTimeFormats bool
var showCustomEntities bool
var showLocalDefaults bool
var showGlobalDefaults bool

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Will show time formats, custom text entity definitions, or defaults.",
	Long:  "Will show time formats, custom text entity definitions, or defaults.",
	Run: func(cmd *cobra.Command, args []string) {
		if !showTimeFormats && !showCustomEntities && !showLocalDefaults && !showGlobalDefaults {
			_ = cmd.Help()
			return
		}

		err := helpers.LoadOutputPrinter()
		if err != nil {
			if helpers.ExitCode == helpers.ExitCodeSuccess {
				// ExitCode was not set in LoadOutputPrinter(), so use general exit code here
				helpers.ExitCode = helpers.ExitCodeUnknownErrorInRootCommand
			}
			if !helpers.CmdHelpers.OutputValueOnly {
				// we use a standard print func here, because the output printer is not available
				fmt.Printf("Critical error in LoadOutputPrinter(): %s\n", err)
			}

			return
		}

		defer func() {
			// Todo: Do something with this error eventually
			_ = helpers.OP.UnloadOutputPrinter()
		}()

		if showTimeFormats {
			printTimeFormats()
		}

		if showCustomEntities {
			printCustomEntities()
		}

		if showLocalDefaults {
			printLocalDefaults()
		}

		if showGlobalDefaults {
			printGlobalDefaults()
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&showTimeFormats, "time-formats", "f", false, "Will show a list of the available time formats")
	showCmd.Flags().BoolVarP(&showCustomEntities, "custom-entities", "c", false, "Will show a list of the available custom text entities")
	showCmd.Flags().BoolVarP(&showLocalDefaults, "local-defaults", "l", false, "Will show the YAML for local default values")
	showCmd.Flags().BoolVarP(&showGlobalDefaults, "global-defaults", "g", false, "Will show the YAML for global default values")
}

func printTimeFormats() {
	helpers.OP.Print(helpers.OutputMode_Force, `
Note: Time Format names are NOT case sensitive.

  Time Formats

  Name               Layout
  =================  ========================================================================
  ANSIC              "Mon Jan _2 15:04:05 2006"
  UnixDate           "Mon Jan _2 15:04:05 MST 2006"
  RubyDate           "Mon Jan 02 15:04:05 -0700 2006"
  RFC822             "02 Jan 06 15:04 MST"
  RFC822Z            "02 Jan 06 15:04 -0700" 
  RFC850             "Monday, 02-Jan-06 15:04:05 MST"
  RFC1123            "Mon, 02 Jan 2006 15:04:05 MST"
  RFC1123Z           "Mon, 02 Jan 2006 15:04:05 -0700" 
  RFC3339            "2006-01-02T15:04:05Z07:00"
  RFC3339Nano        "2006-01-02T15:04:05.999999999Z07:00"
  Kitchen            "3:04PM"
  Stamp              "Jan _2 15:04:05"
  StampMilli         "Jan _2 15:04:05.000"
  StampMicro         "Jan _2 15:04:05.000000"
  StampNano          "Jan _2 15:04:05.000000000"
  USDateTime         "2006-01-02 15:04:05"
  USDateTime         "2006-01-02 15:04:05"
  USDateTimeZ        "2006-01-02 15:04:05 -0700"
  USDateTimeMilliZ   "2006-01-02 15:04:05.000 -0700"
  USDateTimeMicroZ   "2006-01-02 15:04:05.000000 -0700"
  USDateTimeNanoZ    "2006-01-02 15:04:05.000000000 -0700"
  USDateShort        "1/2/06"
  USDate             "01/02/2006"
  EUDateTime         "2006-02-01 15:04:05"
  EUDateTimeZ        "2006-02-01 15:04:05 -0700"
  EUDateTimeMilliZ   "2006-02-01 15:04:05.000 -0700"
  EUDateTimeMicroZ   "2006-02-01 15:04:05.000000 -0700"
  EUDateTimeNanoZ    "2006-02-01 15:04:05.000000000 -0700"
  EUDateShort        "2/1/06"
  EUDate             "02/01/2006"
  DateOnly           "2006-01-02"
  TimeOnly           "15:04:05"
  UnixSecs           Unix Time in seconds
  UnixMilli          Unix Time in millisecond
  UnixMicro          Unix Time in microseconds
  UnixNano           Unix Time in nanoseconds
  Custom             Provide layout text using the flags "--output-layout" and "input-layout" in Timeconverter's formatting syntax
  CustomGO           Provide layout text using the flags "--output-layout" and "input-layout" in Go's fmt formatting syntax
                     See https://pkg.go.dev/time#pkg-constants
`)
}

func printCustomEntities() {
	helpers.OP.Print(helpers.OutputMode_Force, `
Note: Entities references are NOT case sensitive.

  TimeConverter Custom Text Entities Descriptions

  Type                 Description
  ===============      ==================================================================================`)

	for _, desc := range helpers.EntityToDescription {
		helpers.OP.Print(helpers.OutputMode_Force, fmt.Sprintf(
			"  %-15s      %s", desc.Name, desc.Description))
	}

	helpers.OP.Print(helpers.OutputMode_Force, "")
}

func printLocalDefaults() {
	exists, content, err := helpers.CmdHelpers.GetLocalConfigDataIfExists()
	if err != nil {
		helpers.OP.Printf(helpers.OutputMode_Force, "Error reading local defaults: %s", err)
		return
	}

	if !exists {
		helpers.OP.Println(helpers.OutputMode_Force, "Local defaults are not currently set")
		return
	}

	helpers.OP.Print(helpers.OutputMode_Force, `
Local Config Data
============================================================`)

	helpers.OP.Printf(helpers.OutputMode_Force, "%s\n", string(content))
}

func printGlobalDefaults() {
	exists, content, err := helpers.CmdHelpers.GetGlobalConfigDataIfExists()
	if err != nil {
		helpers.OP.Printf(helpers.OutputMode_Force, "Error reading global defaults: %s", err)
		return
	}

	if !exists {
		helpers.OP.Println(helpers.OutputMode_Force, "Global defaults are not currently set")
		return
	}

	helpers.OP.Print(helpers.OutputMode_Force, `
Global Config Data
============================================================`)

	helpers.OP.Printf(helpers.OutputMode_Force, "%s\n", string(content))
}
