// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// IsOffsetTZ examines a timezone reference and determines if it is a local timezone reference
// in IANA region/city format, or if it is an offset ref like -0700, +0000 or +0800.
// Offsets MUST be use a four digit format for the offset value and start with "a" + or "-".
func IsOffsetTZ(tzText string) bool {
	if !strings.Contains(tzText, "/") && len(tzText) == 5 && (tzText[0] == '+' || tzText[0] == '-') {
		return true
	}

	return false
}

// BuildFixedLoc expects input should be in the form +HHMM, where "+" can be "+" or "-", but
// one of them must be supplied.  For UTC, it would be +0000.
// Note that this does not support seconds offsets, just whole hours and mins.
func BuildFixedLoc(tzOffSet string) (loc *time.Location, err error) {
	if len(tzOffSet) != 5 {
		return nil, fmt.Errorf("Incorrect format \"%s\". Expected +HHMM", tzOffSet)
	}

	var signVal int
	switch tzOffSet[0] {
	case '+':
		signVal = 1
	case '-':
		signVal = -1
	default:
		return nil, fmt.Errorf("Incorrect sign indicator \"%s\". Expected \"+\" or \"-\"", string(tzOffSet[0]))
	}

	hoursText := tzOffSet[1:3]
	hours, err := strconv.Atoi(hoursText)
	if err != nil {
		return nil, fmt.Errorf("Hours is not a valid integer: %s", hoursText)
	}

	if hours < 0 || hours > 23 {
		return nil, fmt.Errorf("Hours is out of range: %d.  Must be between 00 and 23", hours)
	}

	minsText := tzOffSet[3:]
	mins, err := strconv.Atoi(minsText)
	if err != nil {
		return nil, fmt.Errorf("Minutes is not a valid integer: %s", minsText)
	}

	if mins < 0 || mins > 23 {
		return nil, fmt.Errorf("Minutes is out of range: %d.  Must be between 00 and 59", mins)
	}

	secondsOffset := signVal * ((hours * 60 * 60) + (mins * 60))

	return time.FixedZone(tzOffSet, secondsOffset), nil
}

// Todo: Maybe add support for abbreviations, perhaps ability to set defaults that indicate which abbreviation to use for ambiguous ones.

// AdjustForOutputTimeZone receives a base time value, then checks to see if the
// user entered value for output-timezone id one of two timezone types.  It then
// adjusts the base time according to the determined offset.
// The timezone construction can be one of the following:
//   - An offset in form of +HHMM
//   - An IANA timezone name in form region/location.
func AdjustForOutputTimeZone(baseTime time.Time) (adjustedTime time.Time, err error) {
	if IsOffsetTZ(CmdHelpers.OutputTimeZone) {
		var fixedZone *time.Location
		fixedZone, err = BuildFixedLoc(CmdHelpers.OutputTimeZone)
		if err != nil {
			return time.Time{}, fmt.Errorf(
				"Unable to parse supplied tzOffset: %s. Error: %s",
				CmdHelpers.OutputTimeZone,
				err)
		}
		return baseTime.In(fixedZone), nil
	}

	// not a tz abbrev, nor a tz offset.  Must be an IANA timezone then.
	tzLoc, err := time.LoadLocation(CmdHelpers.OutputTimeZone)
	if err != nil {
		return time.Time{}, fmt.Errorf("Unable to load indicated output timezone. Error: %s", err)
	}

	return baseTime.In(tzLoc), nil
}
