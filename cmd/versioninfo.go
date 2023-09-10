// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package cmd

import (
	"fmt"
)

// Some of these global vars are placeholders for the build command, when
// relevant values are inserted.

var (
	AppName         = "Timeconverter - A time conversion utility"
	AppMajorVersion = "0"
	AppMinorVersion = "1"
	AppPatchVersion = "2"

	// For any pre-release version, if would need to provide leading ".", like ".dev01"
	AppPreReleaseVer  = ""
	AppVersion        = AppMajorVersion + "." + AppMinorVersion + "." + AppPatchVersion + AppPreReleaseVer
	AppShortBuildTime = "[sbt]"
	AppLongBuildTime  = "[lbt]"
	AppLicense        = "MIT License - Hoby Smith - hoby@thoughtrealm.com"
)

func printVersionInfo(inPromptMode bool) {
	fmt.Println("")
	fmt.Printf("%s\n\n", AppName)
	fmt.Printf("Version          : %s\n", AppVersion)
	fmt.Printf("Build Time[short]: %s\n", AppShortBuildTime)
	fmt.Printf("Build Time[long] : %s\n", AppLongBuildTime)
	fmt.Printf("License          : %s\n", AppLicense)
	fmt.Println("")
}
