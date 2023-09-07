// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package helpers

import (
	"bytes"
	"fmt"
	"golang.design/x/clipboard"
	"os"
	"strings"
	"sync"
)

// OutputPrinter provides an internal wrapper that supports a few simple output mechanisms.
// It provides common output funcs like Print and Printf.  Supported types are Console, Clipboard and Stdout.
type OutputPrinter struct {
	outputTarget    OutputTarget
	outputSync      sync.Mutex
	clipboardBuffer *bytes.Buffer
	file            *os.File
}

// OP is the custom output printer that will send output to the configured target output type based on output mode
var OP *OutputPrinter

func LoadOutputPrinter() error {
	// This routine is called before anything else, so nothing is set yet on cmdHelpers info.
	// We must determine the appropriate output type and then set this logger as needed.

	// Create the basic output printer with defaults.  Really, these are the zero values for
	// these types, but I just set them explicitly here for clarity.
	newOutputPrinter := &OutputPrinter{
		outputTarget: OutputTarget_Console,
		outputSync:   sync.Mutex{},
	}

	// Initialize the output printer
	err := newOutputPrinter.Initialize()
	if err != nil {
		return err
	}

	OP = newOutputPrinter
	return nil
}

func (op *OutputPrinter) UnloadOutputPrinter() (err error) {
	defer func() {
		if r := recover(); r != nil {
			ExitCode = ExitCodePanicInUnloadOutputPrinter
			err = fmt.Errorf("Panic recovered in UnloadOutputPrinter: %s", r)
		}
	}()

	if op.outputTarget == OutputTarget_Clipboard && op.clipboardBuffer != nil {
		err = InitClipboard()
		if err != nil {

		}

		clipboard.Write(clipboard.FmtText, op.clipboardBuffer.Bytes())
	}

	return nil
}

func (op *OutputPrinter) Initialize() error {
	err := op.DetermineOutputType()
	if err != nil {
		ExitCode = ExitCodeErrorDuringInitializeOutputPrinter
		return err
	}

	switch op.outputTarget {
	case OutputTarget_Clipboard:
		op.clipboardBuffer = new(bytes.Buffer)
	default:
		// This can be used for piping chains... for now, this SHOULD work
		op.file = os.Stdout
	}

	return nil
}

func (op *OutputPrinter) DetermineOutputType() error {
	// For thread safety, this "class" owns its own reference to several state vars that are in
	// the CmdHelpers struct.  So, they are both set below, which appears redundant.

	// Determine the type now... first as default, then check for any override behavior needed based on other flags
	CmdHelpers.OutputTarget = OutputTarget_Console
	if CmdHelpers.OutputTargetName != "" {
		var found bool
		CmdHelpers.OutputTarget, found = OutputTargetNameToType[strings.ToLower(CmdHelpers.OutputTargetName)]
		if !found {
			ExitCode = ExitCodeUnknownOutputTargetName
			return fmt.Errorf("Unknown output target name: %s", CmdHelpers.OutputTargetName)
		}
	}
	op.outputTarget = CmdHelpers.OutputTarget

	if CmdHelpers.PipeMode {
		if CmdHelpers.OutputTarget != OutputTarget_Clipboard {
			CmdHelpers.OutputTarget = OutputTarget_Console
		}
	}

	return nil
}

func (op *OutputPrinter) sendOutputText(outputMode OutputMode, outputText string) {
	op.outputSync.Lock()
	defer op.outputSync.Unlock()

	if (CmdHelpers.OutputValueOnly || CmdHelpers.PipeMode) && outputMode != OutputMode_Force {
		return
	}

	// Todo: Handle these errors better
	switch op.outputTarget {
	case OutputTarget_Clipboard:
		op.clipboardBuffer.WriteString(outputText)
	default:
		// Todo: For now, use stdout for default.  It should work fine for the standard print to screen functionality,
		// but if we see issues, we'll deal with it at that point
		_, _ = os.Stdout.WriteString(outputText)
	}
}

func (op *OutputPrinter) Print(outputMode OutputMode, a ...any) {
	op.sendOutputText(outputMode, fmt.Sprintln(a...))
}

func (op *OutputPrinter) Println(outputMode OutputMode, a ...any) {
	op.sendOutputText(outputMode, fmt.Sprintln(a...))
}

func (op *OutputPrinter) Printf(outputMode OutputMode, format string, a ...any) {
	op.sendOutputText(outputMode, fmt.Sprintf(format, a...))
}
