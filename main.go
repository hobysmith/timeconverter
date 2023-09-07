// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package main

import (
	"github.com/hobysmith/timeconverter/cmd"
	"github.com/hobysmith/timeconverter/helpers"
	"os"
	_ "time/tzdata" // This pulls in the IANA TZ database file into the runtime
)

func main() {
	defer func() {
		os.Exit(helpers.ExitCode)
	}()

	cmd.Execute()
}
