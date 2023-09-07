// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

// Package helpers provides all the accessory functionality required by the core and
// command line functionality.  It covers common needs like file wrappers,
// definition of shared constants and types.
// One critical function is providing the structure and behavior around the type HelpersInfo,
// which is used throughout the app for various needs.
package helpers

import (
	"fmt"
	"golang.design/x/clipboard"
)

// HelpersInfo is used to store the command line parameters, as well as transformed values.
// These are set by the cmd package/cobra, and are used or transformed in the primary logic in the encode package.
type HelpersInfo struct {
	// If root.Execute() or maybe other routines want to indicate the causal error, they
	// will set this errResult field.
	ErrResult error
	// String-based input as the source
	Value string `yaml:"-"`
	// The input format name
	InputFormatName string `yaml:"inputFormatName"`
	// The input format type, determined from the input format name
	InputFormat TimeFormat `yaml:"-"`
	// For custom output type, the text of the custom layout
	InputLayout string `yaml:"inputLayout"`
	// The mapped name of the output format
	OutputFormatName string `yaml:"outputFormatName"`
	// The actual TimeFormat type, determined from the OutputFormatName
	OutputFormat TimeFormat `yaml:"-"`
	// For custom output type, the text of the custom layout
	OutputLayout string `yaml:"outputLayout"`
	// The mapped name of the output target
	OutputTargetName string `yaml:"outputTargetName"`
	// The transformed OutputTarget type based on the OutputTargetName
	OutputTarget OutputTarget `yaml:"-"`
	// If true, no output will be emitted
	OutputValueOnly bool `yaml:"outputValueOnly"`
	// The decoded stream
	ConvertedResult string `yaml:"-"`
	// When PipeMode is true, data is read from stdin and written to stdout.
	// Only the converted date is emitted, with possible exceptions for critical errors
	PipeMode bool `yaml:"-"`
	// When SetDefault is true, the currently provided cmd details will be saved as a local path default
	// for future runs.  Local path vals override global vals when both are set.
	SetDefault bool `yaml:"-"`
	// When SetGlobalDefault is true, the currently provided cmd details will be saved as a global
	// path default for future runs.
	SetGlobalDefault bool `yaml:"-"`
	// A timezone to use when converting the output time.  If not specified, the local time will be used for
	// the output time.
	OutputTimeZone string `yaml:"outputTimeZone"`
}

// YamlConfig is used to write out default structures to local and global default files.
type YamlConfig struct {
	ConfigInfo *HelpersInfo `yaml:"TimeConverterSettings"`
}

var ClipboardInitialized bool
var CmdHelpers = &HelpersInfo{}

func InitClipboard() error {
	if !ClipboardInitialized {
		err := clipboard.Init()
		if err != nil {
			return err
		}

		ClipboardInitialized = true
	}

	return nil
}

// InputFormatDesc is called by error handlers to build more descriptive input format names
func (hi *HelpersInfo) InputFormatDesc() string {
	switch hi.InputFormat {
	case TimeFormat_CustomGO:
		return fmt.Sprintf(`CustomGo["%s"]`, hi.InputLayout)
	case TimeFormat_Custom:
		return fmt.Sprintf(`Custom["%s"]`, hi.InputLayout)
	default:
		return TimeFormatToName[hi.InputFormat]
	}
}
