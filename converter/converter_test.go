// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package converter

import (
	_ "time/tzdata"
)

import (
	"fmt"
	"github.com/hobysmith/timeconverter/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Todo: Add separate and more extensive tests for the formats Custom and CustomGo?

// For a test val, we will use an arbitrary date and time of May 7, 2011 14:15:16 EST/-0400

// TestTimeConverter_Convert_Iterate_InputFormats
//
// This test func iterates through all the supported time formats
// and validates that they read input values correctly for those formats.
//
// I do not have specific tests for the formats "Custom" and "CustomGo". They are used a few times in
// these tests, AND I did some separate manual tests against them, so they SHOULD be ok.
// However, I might add specific tests for those formats when I have time.
func TestTimeConverter_Convert_Iterate_InputFormats(t *testing.T) {
	type testInput struct {
		name             string
		inputFormatName  string
		inputLayout      string
		outputFormatName string
		outputLayout     string
		outputTimezone   string
		testInputValue   string
		wantOutputValue  string
		wantErrString    string
	}

	tests := []testInput{
		{
			// "Mon Jan _2 15:04:05 2006"
			name:             "ANSIC",
			inputFormatName:  "ANSIC",
			outputFormatName: "UsDateTime",
			testInputValue:   "Sat May  7 14:15:16 2011",
			wantOutputValue:  "2011-05-07 14:15:16",
		},
		{
			// "Mon Jan _2 15:04:05 MST 2006"
			name:             "UnixDate",
			inputFormatName:  "UnixDate",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "Sat May  7 14:15:16 CDT 2011",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "Mon Jan 02 15:04:05 -0700 2006"
			name:             "RubyDate",
			inputFormatName:  "RubyDate",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "Sat May 07 14:15:16 -0500 2011",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "02 Jan 06 15:04 MST"
			name:             "RFC822",
			inputFormatName:  "RFC822",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "07 May 11 14:15 EST",
			wantOutputValue:  "2011-05-07 14:15:00 -0500", // this format does not respect century date values
		},
		{
			// "02 Jan 06 15:04 -0700"
			name:             "RFC822Z",
			inputFormatName:  "RFC822Z",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "07 May 11 14:15 -0500",
			wantOutputValue:  "2011-05-07 14:15:00 -0500",
		},
		{
			// "Monday, 02-Jan-06 15:04:05 MST"
			name:             "RFC850",
			inputFormatName:  "RFC850",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "Saturday, 07-May-11 14:15:16 EST",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "Mon, 02 Jan 2006 15:04:05 MST"
			name:             "RFC1123",
			inputFormatName:  "RFC1123",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "Sat, 07 May 2011 14:15:16 EST",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "Mon, 02 Jan 2006 15:04:05 -0700"
			name:             "RFC1123Z",
			inputFormatName:  "RFC1123Z",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "Sat, 07 May 2011 14:15:16 -0500",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "2006-01-02T15:04:05Z07:00"
			name:             "RFC3339",
			inputFormatName:  "RFC3339",
			outputFormatName: "UsDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07T14:15:16-05:00",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "2006-01-02T15:04:05.999999999Z07:00"
			name:             "RFC3339Nano",
			inputFormatName:  "RFC3339Nano",
			outputFormatName: "UsDateTimeNanoZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07T14:15:16.123456789-05:00",
			wantOutputValue:  "2011-05-07 14:15:16.123456789 -0500",
		},
		{
			// "3:04PM"
			name:             "Kitchen",
			inputFormatName:  "Kitchen",
			outputFormatName: "TimeOnly",
			testInputValue:   "2:15PM",
			wantOutputValue:  "14:15:00",
		},
		{
			// "Jan _2 15:04:05"
			name:             "Stamp",
			inputFormatName:  "Stamp",
			outputFormatName: "Custom",
			outputLayout:     "mmm d hhh:nn:ss",
			testInputValue:   "May 7 14:15:16",
			wantOutputValue:  "May 7 14:15:16",
		},
		{
			// "Jan _2 15:04:05.000"
			name:             "StampMilli",
			inputFormatName:  "StampMilli",
			outputFormatName: "Custom",
			outputLayout:     "mmm d hhh:nn:ss.zzz",
			testInputValue:   "May  7 14:15:16.123",
			wantOutputValue:  "May 7 14:15:16.123",
		},
		{
			// "Jan _2 15:04:05.000000"
			name:             "StampMicro",
			inputFormatName:  "StampMicro",
			outputFormatName: "Custom",
			outputLayout:     "mmm d hhh:nn:ss.zzzzzz",
			testInputValue:   "May  7 14:15:16.123456",
			wantOutputValue:  "May 7 14:15:16.123456",
		},
		{
			// "Jan _2 15:04:05.000000000"
			name:             "StampNano",
			inputFormatName:  "StampNano",
			outputFormatName: "Custom",
			outputLayout:     "mmm d hhh:nn:ss.zzzzzzzzz",
			testInputValue:   "May  7 14:15:16.123456789",
			wantOutputValue:  "May 7 14:15:16.123456789",
		},
		{
			// "2006-01-02 15:04:05"
			name:             "USDateTime",
			inputFormatName:  "USDateTime",
			outputFormatName: "USDateTime",
			testInputValue:   "2011-05-07 14:15:16",
			wantOutputValue:  "2011-05-07 14:15:16",
		},
		{
			// "2006-01-02 15:04:05"
			name:             "USDateTimeZ",
			inputFormatName:  "USDateTimeZ",
			outputFormatName: "USDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "2006-01-02 15:04:05.000 -0700"
			name:             "USDateTimeMilliZ",
			inputFormatName:  "USDateTimeMilliZ",
			outputFormatName: "USDateTimeMilliZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123 -0500",
		},
		{
			// "2006-01-02 15:04:05.000000 -0700"
			name:             "USDateTimeMicroZ",
			inputFormatName:  "USDateTimeMicroZ",
			outputFormatName: "USDateTimeMicroZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123456 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123456 -0500",
		},
		{
			// "2006-01-02 15:04:05.000000000 -0700"
			name:             "USDateTimeNanoZ",
			inputFormatName:  "USDateTimeNanoZ",
			outputFormatName: "USDateTimeNanoZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123456789 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123456789 -0500",
		},
		{
			// "1/2/06"
			name:             "USDateShort",
			inputFormatName:  "USDateShort",
			outputFormatName: "USDate",
			testInputValue:   "5/7/11",
			wantOutputValue:  "05/07/2011",
		},
		{
			// "01/02/2006"
			name:             "USDate",
			inputFormatName:  "USDate",
			outputFormatName: "USDateTime",
			testInputValue:   "05/07/2011",
			wantOutputValue:  "2011-05-07 00:00:00",
		},
		{
			// "2006-02-01 15:04:05"
			name:             "EUDateTime",
			inputFormatName:  "EUDateTime",
			outputFormatName: "USDateTime",
			testInputValue:   "2011-07-05 14:15:16",
			wantOutputValue:  "2011-05-07 14:15:16",
		},
		{
			// "2006-02-01 15:04:05 -0700"
			name:             "EUDateTimeZ",
			inputFormatName:  "EUDateTimeZ",
			outputFormatName: "USDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-07-05 14:15:16 -0500",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "2006-02-01 15:04:05.000 -0700"
			name:             "EUDateTimeMilliZ",
			inputFormatName:  "EUDateTimeMilliZ",
			outputFormatName: "USDateTimeMilliZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-07-05 14:15:16.123 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123 -0500",
		},
		{
			// "2006-02-01 15:04:05.000000 -0700"
			name:             "EUDateTimeMicroZ",
			inputFormatName:  "EUDateTimeMicroZ",
			outputFormatName: "USDateTimeMicroZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-07-05 14:15:16.123456 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123456 -0500",
		},
		{
			// "2006-02-01 15:04:05.000000000 -0700"
			name:             "EUDateTimeNanoZ",
			inputFormatName:  "EUDateTimeNanoZ",
			outputFormatName: "USDateTimeNanoZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-07-05 14:15:16.123456789 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123456789 -0500",
		},
		{
			// "2006-02-01 15:04:05.000000000 -0700"
			name:             "EUDateShort",
			inputFormatName:  "EUDateShort",
			outputFormatName: "USDateTime",
			testInputValue:   "7/5/11",
			wantOutputValue:  "2011-05-07 00:00:00",
		},
		{
			// "2/1/06"
			name:             "EUDateShort",
			inputFormatName:  "EUDateShort",
			outputFormatName: "USDateTime",
			testInputValue:   "7/5/11",
			wantOutputValue:  "2011-05-07 00:00:00",
		},
		{
			// "02/01/2006"
			name:             "EUDate",
			inputFormatName:  "EUDate",
			outputFormatName: "USDate",
			testInputValue:   "07/05/2011",
			wantOutputValue:  "05/07/2011",
		},
		{
			// "2006-01-02"
			name:             "DateOnly",
			inputFormatName:  "DateOnly",
			outputFormatName: "USDateTime",
			testInputValue:   "2011-05-07",
			wantOutputValue:  "2011-05-07 00:00:00",
		},
		{
			// "15:04:05"
			name:             "TimeOnly",
			inputFormatName:  "TimeOnly",
			outputFormatName: "CustomGo",
			outputLayout:     "15:04:05",
			testInputValue:   "14:15:16",
			wantOutputValue:  "14:15:16",
		},

		// The following are specific cases or edge cases that
		// are not covered by the defs above
		{
			// when inputformat is not defined, it should default to USDateTimeZ
			name:             "InputFormatNotDefined",
			inputFormatName:  "",
			outputFormatName: "USDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// when input value is not provided, an error should be returned
			name:             "InputValueNotProvided",
			inputFormatName:  "USDateTimeZ",
			outputFormatName: "USDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
			wantErrString:    "No input provided",
		},
	}

	for idx, test := range tests {
		t.Run(
			fmt.Sprintf(
				"[%02d] "+test.name+`: input value "%s"`,
				idx+1,
				test.testInputValue,
			),
			func(t *testing.T) {
				helpers.CmdHelpers.InputFormatName = test.inputFormatName
				helpers.CmdHelpers.InputLayout = test.inputLayout
				helpers.CmdHelpers.Value = test.testInputValue
				helpers.CmdHelpers.OutputFormatName = test.outputFormatName
				helpers.CmdHelpers.OutputLayout = test.outputLayout
				helpers.CmdHelpers.OutputTimeZone = test.outputTimezone

				err := New().Convert(true)
				if test.wantErrString != "" {
					assert.Contains(t, err.Error(), test.wantErrString)
				} else {
					assert.Nil(t, err)
				}
			},
		)
	}
}

// TestTimeConverter_Convert_Iterate_OutputFormats iterates through all the supported time formats
// and validates that they write output values correctly for those formats.
//
// For further details, see the inline notes for TestTimeConverter_Convert_Iterate_OutputFormats.
func TestTimeConverter_Convert_Iterate_OutputFormats(t *testing.T) {
	type testOutput struct {
		name             string
		inputFormatName  string
		inputLayout      string
		outputFormatName string
		outputLayout     string
		outputTimezone   string
		testInputValue   string
		wantOutputValue  string
		wantErrString    string
	}

	tests := []testOutput{
		{
			// "Mon Jan _2 15:04:05 2006"
			name:             "ANSIC",
			inputFormatName:  "USDateTime",
			outputFormatName: "ANSIC",
			testInputValue:   "2011-05-07 14:15:16",
			wantOutputValue:  "Sat May  7 14:15:16 2011",
		},
		{
			// "Mon Jan _2 15:04:05 MST 2006"
			name:             "UnixDate",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "UnixDate",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "Sat May  7 14:15:16 CDT 2011",
		},
		{
			// "Mon Jan 02 15:04:05 -0700 2006"
			name:             "RubyDate",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "RubyDate",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "Sat May 07 14:15:16 -0500 2011",
		},
		{
			// "02 Jan 06 15:04 MST"
			name:             "RFC822",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "RFC822",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:00 -0500",
			wantOutputValue:  "07 May 11 14:15 CDT", // this format does not respect century date values
		},
		{
			// "02 Jan 06 15:04 -0700"
			name:             "RFC822Z",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "RFC822Z",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:00 -0500",
			wantOutputValue:  "07 May 11 14:15 -0500",
		},
		{
			// "Monday, 02-Jan-06 15:04:05 MST"
			name:             "RFC850",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "RFC850",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "Saturday, 07-May-11 14:15:16 CDT",
		},
		{
			// "Mon, 02 Jan 2006 15:04:05 MST"
			name:             "RFC1123",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "RFC1123",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "Sat, 07 May 2011 14:15:16 CDT",
		},
		{
			// "Mon, 02 Jan 2006 15:04:05 -0700",
			name:             "RFC1123Z",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "RFC1123Z",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "Sat, 07 May 2011 14:15:16 -0500",
		},
		{
			// "2006-01-02T15:04:05Z07:00",
			name:             "RFC3339",
			inputFormatName:  "UsDateTimeZ",
			outputFormatName: "RFC3339",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "2011-05-07T14:15:16-05:00",
		},
		{
			// "2006-01-02T15:04:05.999999999Z07:00",
			name:             "RFC3339Nano",
			inputFormatName:  "UsDateTimeNanoZ",
			outputFormatName: "RFC3339Nano",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123456789 -0500",
			wantOutputValue:  "2011-05-07T14:15:16.123456789-05:00",
		},
		{
			// "3:04PM",
			name:             "Kitchen",
			inputFormatName:  "TimeOnly",
			outputFormatName: "Kitchen",
			testInputValue:   "14:15:00",
			wantOutputValue:  "2:15PM",
		},
		{
			// "Jan _2 15:04:05",
			name:             "Stamp",
			inputFormatName:  "Custom",
			inputLayout:      "mmm d hhh:nn:ss",
			outputFormatName: "Stamp",
			testInputValue:   "May 7 14:15:16",
			wantOutputValue:  "May  7 14:15:16",
		},
		{
			// "Jan _2 15:04:05.000",
			name:             "StampMilli",
			inputFormatName:  "Custom",
			inputLayout:      "mmm d hhh:nn:ss.zzz",
			outputFormatName: "StampMilli",
			testInputValue:   "May 7 14:15:16.123",
			wantOutputValue:  "May  7 14:15:16.123",
		},
		{
			// "Jan _2 15:04:05.000000",
			name:             "StampMicro",
			inputFormatName:  "Custom",
			inputLayout:      "mmm d hhh:nn:ss.zzzzzz",
			outputFormatName: "StampMicro",
			testInputValue:   "May 7 14:15:16.123456",
			wantOutputValue:  "May  7 14:15:16.123456",
		},
		{
			// "Jan _2 15:04:05.000000000",
			name:             "StampNano",
			inputFormatName:  "Custom",
			inputLayout:      "mmm d hhh:nn:ss.zzzzzzzzz",
			outputFormatName: "StampNano",
			testInputValue:   "May 7 14:15:16.123456789",
			wantOutputValue:  "May  7 14:15:16.123456789",
		},
		{
			// "2006-01-02 15:04:05",
			name:             "USDateTime",
			inputFormatName:  "USDateTime",
			outputFormatName: "USDateTime",
			testInputValue:   "2011-05-07 14:15:16",
			wantOutputValue:  "2011-05-07 14:15:16",
		},
		{
			// "2006-01-02 15:04:05",
			name:             "USDateTimeZ",
			inputFormatName:  "USDateTimeZ",
			outputFormatName: "USDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
		{
			// "2006-01-02 15:04:05.000 -0700",
			name:             "USDateTimeMilliZ",
			inputFormatName:  "USDateTimeMilliZ",
			outputFormatName: "USDateTimeMilliZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123 -0500",
		},
		{
			// "2006-01-02 15:04:05.000000 -0700",
			name:             "USDateTimeMicroZ",
			inputFormatName:  "USDateTimeMicroZ",
			outputFormatName: "USDateTimeMicroZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123456 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123456 -0500",
		},
		{
			// "2006-01-02 15:04:05.000000000 -0700",
			name:             "USDateTimeNanoZ",
			inputFormatName:  "USDateTimeNanoZ",
			outputFormatName: "USDateTimeNanoZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123456789 -0500",
			wantOutputValue:  "2011-05-07 14:15:16.123456789 -0500",
		},
		{
			// "1/2/06",
			name:             "USDateShort",
			inputFormatName:  "USDate",
			outputFormatName: "USDateShort",
			testInputValue:   "05/07/2011",
			wantOutputValue:  "5/7/11",
		},
		{
			// "01/02/2006",
			name:             "USDate",
			inputFormatName:  "USDateTime",
			outputFormatName: "USDate",
			testInputValue:   "2011-05-07 00:00:00",
			wantOutputValue:  "05/07/2011",
		},
		{
			// "2006-02-01 15:04:05",
			name:             "EUDateTime",
			inputFormatName:  "USDateTime",
			outputFormatName: "EUDateTime",
			testInputValue:   "2011-05-07 14:15:16",
			wantOutputValue:  "2011-07-05 14:15:16",
		},
		{
			// "2006-02-01 15:04:05 -0700",
			name:             "EUDateTimeZ",
			inputFormatName:  "USDateTimeZ",
			outputFormatName: "EUDateTimeZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "2011-07-05 14:15:16 -0500",
		},
		{
			// "2006-02-01 15:04:05.000 -0700",
			name:             "EUDateTimeMilliZ",
			inputFormatName:  "USDateTimeMilliZ",
			outputFormatName: "EUDateTimeMilliZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123 -0500",
			wantOutputValue:  "2011-07-05 14:15:16.123 -0500",
		},
		{
			// "2006-02-01 15:04:05.000000 -0700",
			name:             "EUDateTimeMicroZ",
			inputFormatName:  "USDateTimeMicroZ",
			outputFormatName: "EUDateTimeMicroZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123456 -0500",
			wantOutputValue:  "2011-07-05 14:15:16.123456 -0500",
		},
		{
			// "2006-02-01 15:04:05.000000000 -0700",
			name:             "EUDateTimeNanoZ",
			inputFormatName:  "USDateTimeNanoZ",
			outputFormatName: "EUDateTimeNanoZ",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16.123456789 -0500",
			wantOutputValue:  "2011-07-05 14:15:16.123456789 -0500",
		},
		{
			// "2006-02-01 15:04:05.000000000 -0700",
			name:             "EUDateShort",
			inputFormatName:  "USDateTime",
			outputFormatName: "EUDateShort",
			testInputValue:   "2011-05-07 00:00:00",
			wantOutputValue:  "7/5/11",
		},
		{
			// "2/1/06",
			name:             "EUDateShort",
			inputFormatName:  "USDateTime",
			outputFormatName: "EUDateShort",
			testInputValue:   "2011-05-07 00:00:00",
			wantOutputValue:  "7/5/11",
		},
		{
			// "02/01/2006",
			name:             "EUDate",
			inputFormatName:  "USDate",
			outputFormatName: "EUDate",
			testInputValue:   "05/07/2011",
			wantOutputValue:  "07/05/2011",
		},
		{
			// "2006-01-02",
			name:             "DateOnly",
			inputFormatName:  "USDateTime",
			outputFormatName: "DateOnly",
			testInputValue:   "2011-05-07 00:00:00",
			wantOutputValue:  "2011-05-07",
		},
		{
			// "15:04:05",
			name:             "TimeOnly",
			inputFormatName:  "CustomGo",
			inputLayout:      "15:04:05",
			outputFormatName: "TimeOnly",
			testInputValue:   "14:15:16",
			wantOutputValue:  "14:15:16",
		},

		// The following are specific cases or edge cases that
		// are not covered by the defs above
		{
			// when outputformat is not defined, it should default to USDateTimeZ
			name:             "OutputFormatNotDefined",
			inputFormatName:  "USDateTimeZ",
			outputFormatName: "",
			outputTimezone:   "-0500",
			testInputValue:   "2011-05-07 14:15:16 -0500",
			wantOutputValue:  "2011-05-07 14:15:16 -0500",
		},
	}

	for idx, test := range tests {
		t.Run(
			fmt.Sprintf(
				"[%02d] "+test.name+`: input value "%s"`,
				idx+1,
				test.testInputValue,
			),
			func(t *testing.T) {
				helpers.CmdHelpers.InputFormatName = test.inputFormatName
				helpers.CmdHelpers.InputLayout = test.inputLayout
				helpers.CmdHelpers.Value = test.testInputValue
				helpers.CmdHelpers.OutputFormatName = test.outputFormatName
				helpers.CmdHelpers.OutputLayout = test.outputLayout
				helpers.CmdHelpers.OutputTimeZone = test.outputTimezone

				err := New().Convert(true)
				if test.wantErrString != "" {
					assert.Contains(t, err.Error(), test.wantErrString)
				} else {
					assert.Nil(t, err)
				}
			},
		)
	}
}
