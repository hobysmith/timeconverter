// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package helpers

import "os"

// this should check to see if pipe input is available

func CheckIsPiped() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		CmdHelpers.PipeMode = true
	}

	return CmdHelpers.PipeMode
}
