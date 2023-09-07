// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package helpers

import (
	"errors"
	"os"
)

func FileExists(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err == nil {
		// the path ref does exist, but it could be pointing to a directory, which is not what we want
		if info.IsDir() {
			return false, errors.New("specified file path is a directory")
		}

		// not a directory, so this is ok
		return true, nil
	}

	// check if the error was simply due to file not existing
	if errors.Is(err, os.ErrNotExist) {
		// no hard error, file just doesn't exist
		return false, nil
	}

	return false, err
}
