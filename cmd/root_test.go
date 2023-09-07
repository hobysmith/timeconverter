// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package cmd

import (
	"github.com/hobysmith/timeconverter/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This root_test implements some basic validations around the command line
// functionality, such as flag settings, calling interfaces to functional packages, etc.
// There will not be an exhaustive set of tests here.  There are more exhaustive
// tests of the core transformation functionality in the converter_test file.

func TestSimpleConversion_USDateTimeZToUSDateTimeZ(t *testing.T) {
	c := GetRootCmd()
	c.SetArgs([]string{"2023-08-25 23:52:10 -0700", "-i=USDateTimeZ", "-o=USDateTimeZ"})
	err := c.Execute()
	assert.Nil(t, err)
	assert.Equal(t, "2023-08-25 23:52:10 -0700", helpers.CmdHelpers.ConvertedResult)
}

func TestSimpleConversion_FailsOnUnknownInputFormat(t *testing.T) {
	c := GetRootCmd()
	c.SetArgs([]string{"2023-08-25 23:52:10 -0700", "-i=Unknown", "-o=USDateTimeZ", "-v"})
	err := c.Execute()
	assert.Nil(t, err) // not a catastrophic error
	assert.NotNil(t, helpers.CmdHelpers.ErrResult)
	assert.Contains(t, helpers.CmdHelpers.ErrResult.Error(), "Unknown input-format: Unknown")
}

func TestSimpleConversion_FailsOnUnknownOutputFormat(t *testing.T) {
	c := GetRootCmd()
	c.SetArgs([]string{"2023-08-25 23:52:10 -0700", "-i=USDateTimeZ", "-o=Unknown", "-v"})
	err := c.Execute()
	assert.Nil(t, err) // not a catastrophic error
	assert.NotNil(t, helpers.CmdHelpers.ErrResult)
	assert.Contains(t, helpers.CmdHelpers.ErrResult.Error(), "Unknown output-format: Unknown")
}
