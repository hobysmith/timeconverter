// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

// Package converter provides a wrapper for the functionality necessary to
// transform date value from one format to another.  The cmd package parses
// commands, preps relevant settings, then calls the TimeConverter wrapper
// to do the work.
package converter

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hobysmith/timeconverter/helpers"
	"os"
	"strconv"
	"strings"
	"time"
)

// TypeConverter is the primary wrapper for conversion functionality.
type TimeConverter struct{}

func New() *TimeConverter {
	return &TimeConverter{}
}

// Convert executes the main conversion path.
// Note that the parameter fullQuiet means NOTHING should be output from this app,
// which should generally only be used by test funcs.
func (tfd *TimeConverter) Convert(fullQuiet bool) (err error) {
	// Make sure that only 1 input value has been provided
	var inputVal string

	if helpers.CmdHelpers.PipeMode {
		inputValBytes, err := tfd.GetPipeInput()
		if err != nil {
			helpers.ExitCode = helpers.ExitCodeFailureReadingPipeInput
			return fmt.Errorf("Failure reading pipe input: %s", err)
		}
		inputVal = strings.Trim(string(inputValBytes), "\n\r\t ")
	} else {
		inputVal = helpers.CmdHelpers.Value
	}

	if len(inputVal) == 0 {
		helpers.ExitCode = helpers.ExitCodeErrorNoInputProvided
		return errors.New("No input provided")
	}

	if helpers.CmdHelpers.InputFormatName == "" {
		helpers.CmdHelpers.InputFormat = helpers.TimeFormat_USDateTimeZ
	} else {
		found := false
		helpers.CmdHelpers.InputFormat, found = helpers.NameToTimeFormat[strings.ToUpper(helpers.CmdHelpers.InputFormatName)]
		if !found {
			return fmt.Errorf("Unknown input-format: %s", helpers.CmdHelpers.InputFormatName)
		}
	}

	if helpers.CmdHelpers.OutputFormatName == "" {
		helpers.CmdHelpers.OutputFormat = helpers.TimeFormat_USDateTimeZ
	} else {
		found := false
		helpers.CmdHelpers.OutputFormat, found = helpers.NameToTimeFormat[strings.ToUpper(helpers.CmdHelpers.OutputFormatName)]
		if !found {
			return fmt.Errorf("Unknown output-format: %s", helpers.CmdHelpers.OutputFormatName)
		}
	}

	var convertedTime time.Time

	if strings.ToLower(inputVal) == "now" {
		convertedTime = time.Now()
	} else {
		convertedTime, err = tfd.ParseInputTime(inputVal, helpers.CmdHelpers.InputFormat)
		if err != nil {
			return fmt.Errorf(
				"Unable to parse \"%s\" using format %s. Input value or format is not correct.",
				inputVal,
				helpers.CmdHelpers.InputFormatDesc(),
			)
		}
	}

	if helpers.CmdHelpers.OutputTimeZone != "" {
		convertedTime, err = helpers.AdjustForOutputTimeZone(convertedTime)
		if err != nil {
			return err
		}
	}

	helpers.CmdHelpers.ConvertedResult, err = helpers.NewDateTimeFormatter(convertedTime).FormatDateTime(helpers.CmdHelpers.OutputFormat)
	if err != nil {
		return fmt.Errorf("Critical error: Failure converting input to formatted result: %s", err)
	}

	if fullQuiet {
		// fullQuiet means NOTHING should be output, which is generally only used by test funcs
		return
	}

	if helpers.CmdHelpers.OutputValueOnly {
		helpers.OP.Printf(helpers.OutputMode_Force, "%s\n", helpers.CmdHelpers.ConvertedResult)
	} else {
		helpers.OP.Printf(helpers.OutputMode_Force, "Converted Result: %s\n", helpers.CmdHelpers.ConvertedResult)
	}

	return nil
}

// GetPipeInput is called to retrieve data from StdIn
func (tfd *TimeConverter) GetPipeInput() ([]byte, error) {
	pipeBuffer := new(bytes.Buffer)
	_, err := pipeBuffer.ReadFrom(os.Stdin)
	if err != nil {
		return nil, err
	}

	return pipeBuffer.Bytes(), nil
}

// ParseInputTime is called to transform the input value text into a time.Time value based on inputFormat
func (tfd *TimeConverter) ParseInputTime(inputTimeText string, inputFormat helpers.TimeFormat) (convertTime time.Time, err error) {
	var inputUnixInt int64
	if helpers.IsUnixTimeFormat(inputFormat) {
		inputUnixInt, err = strconv.ParseInt(string(inputTimeText), 10, 64)
		if err != nil {
			helpers.ExitCode = helpers.ExitCodeErrorDecodingInput
			return time.Time{}, err
		}
	}

	var layout string
	switch inputFormat {
	case helpers.TimeFormat_Unix_Secs:
		return time.Unix(inputUnixInt, 0), nil
	case helpers.TimeFormat_Unix_Milli:
		return time.UnixMilli(inputUnixInt), nil
	case helpers.TimeFormat_Unix_Micro:
		return time.UnixMicro(inputUnixInt), nil
	case helpers.TimeFormat_Unix_Nano:
		return time.Unix(0, inputUnixInt), nil
	case helpers.TimeFormat_CustomGO:
		layout = helpers.CmdHelpers.InputLayout
	case helpers.TimeFormat_Custom:
		formatter := helpers.NewDateTimeFormatter(time.Now())
		layout, err = formatter.BuildCustomLayout(helpers.CmdHelpers.InputLayout)
		if err != nil {
			helpers.ExitCode = helpers.ExitCodeErrorDecodingInput
			return time.Time{}, err
		}
	default:
		layout = helpers.TimeFormatToLayout[inputFormat]
	}

	return time.Parse(layout, inputTimeText)
}
