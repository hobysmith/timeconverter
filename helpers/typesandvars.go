// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package helpers

// ExitCode is used as the final ExitCode return by this runtime
var ExitCode int = ExitCodeSuccess

// A list of ExitCode value mappings to specific errors.
// This is useful when running this utility from a shell call or some other app.
const (
	ExitCodeSuccess = iota
	ExitCodePanicInExecute
	ExitCodeErrorReturnedToExecute
	ExitCodeUnknownErrorInRootCommand
	ExitCodeUnknownOutputTargetName
	ExitCodeFailureReadingPipeInput
	ExitCodePanicInUnloadOutputPrinter
	ExitCodeErrorDuringInitializeOutputPrinter
	ExitCodeErrorDecodingInput
	ExitCodeErrorNoInputProvided
)

type OutputMode int

const (
	OutputMode_Verbose OutputMode = iota
	OutputMode_Terse
	OutputMode_Force
)

var OutputModeNameToMode = map[string]OutputMode{
	"verbose": OutputMode_Verbose,
	"terse":   OutputMode_Terse,
}

type OutputTarget int

const (
	OutputTarget_Console OutputTarget = iota
	OutputTarget_Clipboard
)

var OutputTargetToName = map[OutputTarget]string{
	OutputTarget_Console:   "console",
	OutputTarget_Clipboard: "clipboard",
}

var OutputTargetNameToType = map[string]OutputTarget{
	"console":   OutputTarget_Console,
	"clipboard": OutputTarget_Clipboard,
}

type TimeFormat int

const (
	TimeFormat_ANSIC            TimeFormat = iota // "Mon Jan _2 15:04:05 2006"
	TimeFormat_UnixDate                           // "Mon Jan _2 15:04:05 MST 2006"
	TimeFormat_RubyDate                           // "Mon Jan 02 15:04:05 -0700 2006"
	TimeFormat_RFC822                             // "02 Jan 06 15:04 MST"
	TimeFormat_RFC822Z                            // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	TimeFormat_RFC850                             // "Monday, 02-Jan-06 15:04:05 MST"
	TimeFormat_RFC1123                            // "Mon, 02 Jan 2006 15:04:05 MST"
	TimeFormat_RFC1123Z                           // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	TimeFormat_RFC3339                            // "2006-01-02T15:04:05Z07:00"
	TimeFormat_RFC3339Nano                        // "2006-01-02T15:04:05.999999999Z07:00"
	TimeFormat_Kitchen                            // "3:04PM"
	TimeFormat_Stamp                              // "Jan _2 15:04:05"
	TimeFormat_StampMilli                         // "Jan _2 15:04:05.000"
	TimeFormat_StampMicro                         // "Jan _2 15:04:05.000000"
	TimeFormat_StampNano                          // "Jan _2 15:04:05.000000000"
	TimeFormat_USDateTime                         // "2006-01-02 15:04:05"
	TimeFormat_USDateTimeZ                        // "2006-01-02 15:04:05 -0700"
	TimeFormat_USDateTimeMilliZ                   // "2006-01-02 15:04:05.000 -0700"
	TimeFormat_USDateTimeMicroZ                   // "2006-01-02 15:04:05.000000 -0700"
	TimeFormat_USDateTimeNanoZ                    // "2006-01-02 15:04:05.000000000 -0700"
	TimeFormat_USDateShort                        // "1/2/06"
	TimeFormat_USDate                             // "01/02/2006"
	TimeFormat_EUDateTime                         // "2006-02-01 15:04:05"
	TimeFormat_EUDateTimeZ                        // "2006-02-01 15:04:05 -0700"
	TimeFormat_EUDateTimeMilliZ                   // "2006-02-01 15:04:05.000 -0700"
	TimeFormat_EUDateTimeMicroZ                   // "2006-02-01 15:04:05.000000 -0700"
	TimeFormat_EUDateTimeNanoZ                    // "2006-02-01 15:04:05.000000000 -0700"
	TimeFormat_DateOnly                           // "2006-01-02"
	TimeFormat_TimeOnly                           // "15:04:05"
	TimeFormat_Custom                             // Specify format with yy/yyyy m/mm/mmm/MMM/Mmm/mmmm/MMMM/Mmmm d/dd/ddd/DDD/Ddd/dddd/DDDD/Dddd h/hh n/nn ss zzz/zzzzzz/zzzzzzzzz tz
	TimeFormat_CustomGO                           // Specify format with GO layout specs
	TimeFormat_EUDateShort                        // "2/1/06"
	TimeFormat_EUDate                             // "02/01/2006"
	TimeFormat_Unix_Secs                          // Unix Seconds
	TimeFormat_Unix_Milli                         // Unix Milliseconds
	TimeFormat_Unix_Micro                         // Unix Microseconds
	TimeFormat_Unix_Nano                          // Unix Nanoseconds
)

var NameToTimeFormat = map[string]TimeFormat{
	"ANSIC":            TimeFormat_ANSIC,
	"UNIXDATE":         TimeFormat_UnixDate,
	"RUBYDATE":         TimeFormat_RubyDate,
	"RFC822":           TimeFormat_RFC822,
	"RFC822Z":          TimeFormat_RFC822Z,
	"RFC850":           TimeFormat_RFC850,
	"RFC1123":          TimeFormat_RFC1123,
	"RFC1123Z":         TimeFormat_RFC1123Z,
	"RFC3339":          TimeFormat_RFC3339,
	"RFC3339NANO":      TimeFormat_RFC3339Nano,
	"KITCHEN":          TimeFormat_Kitchen,
	"STAMP":            TimeFormat_Stamp,
	"STAMPMILLI":       TimeFormat_StampMilli,
	"STAMPMICRO":       TimeFormat_StampMicro,
	"STAMPNANO":        TimeFormat_StampNano,
	"USDATETIME":       TimeFormat_USDateTime,
	"USDATETIMEZ":      TimeFormat_USDateTimeZ,
	"USDATETIMEMILLIZ": TimeFormat_USDateTimeMilliZ,
	"USDATETIMEMICROZ": TimeFormat_USDateTimeMicroZ,
	"USDATETIMENANOZ":  TimeFormat_USDateTimeNanoZ,
	"USDATESHORT":      TimeFormat_USDateShort,
	"USDATE":           TimeFormat_USDate,
	"EUDATETIME":       TimeFormat_EUDateTime,
	"EUDATETIMEZ":      TimeFormat_EUDateTimeZ,
	"EUDATETIMEMILLIZ": TimeFormat_EUDateTimeMilliZ,
	"EUDATETIMEMICROZ": TimeFormat_EUDateTimeMicroZ,
	"EUDATETIMENANOZ":  TimeFormat_EUDateTimeNanoZ,
	"EUDATESHORT":      TimeFormat_EUDateShort,
	"EUDATE":           TimeFormat_EUDate,
	"DATEONLY":         TimeFormat_DateOnly,
	"TIMEONLY":         TimeFormat_TimeOnly,
	"CUSTOM":           TimeFormat_Custom,
	"CUSTOMGO":         TimeFormat_CustomGO,
	"UNIXSECS":         TimeFormat_Unix_Secs,
	"UNIXMILLI":        TimeFormat_Unix_Milli,
	"UNIXMICRO":        TimeFormat_Unix_Micro,
	"UNIXNANO":         TimeFormat_Unix_Nano,
}

var TimeFormatToName = map[TimeFormat]string{
	TimeFormat_ANSIC:            "ANSIC",
	TimeFormat_UnixDate:         "UnixDate",
	TimeFormat_RubyDate:         "RubyDate",
	TimeFormat_RFC822:           "RFC822",
	TimeFormat_RFC822Z:          "RFC822Z",
	TimeFormat_RFC850:           "RFC850",
	TimeFormat_RFC1123:          "RFC1123",
	TimeFormat_RFC1123Z:         "RFC1123Z",
	TimeFormat_RFC3339:          "RFC3339",
	TimeFormat_RFC3339Nano:      "RFC3339Nano",
	TimeFormat_Kitchen:          "Kitchen",
	TimeFormat_Stamp:            "Stamp",
	TimeFormat_StampMilli:       "StampMilli",
	TimeFormat_StampMicro:       "StampMicro",
	TimeFormat_StampNano:        "StampNano",
	TimeFormat_USDateTime:       "USDateTime",
	TimeFormat_USDateTimeZ:      "USDateTimeZ",
	TimeFormat_USDateTimeMilliZ: "USDateTimeMilliZ",
	TimeFormat_USDateTimeMicroZ: "USDateTimeMicroZ",
	TimeFormat_USDateTimeNanoZ:  "USDateTimeNanoZ",
	TimeFormat_USDateShort:      "USDateShort",
	TimeFormat_USDate:           "USDate",
	TimeFormat_EUDateTime:       "EUDateTime",
	TimeFormat_EUDateTimeZ:      "EUDateTimeZ",
	TimeFormat_EUDateTimeMilliZ: "EUDateTimeMilliZ",
	TimeFormat_EUDateTimeMicroZ: "EUDateTimeMicroZ",
	TimeFormat_EUDateTimeNanoZ:  "EUDateTimeNanoZ",
	TimeFormat_EUDateShort:      "EUDateShort",
	TimeFormat_EUDate:           "EUDate",
	TimeFormat_DateOnly:         "DateOnly",
	TimeFormat_TimeOnly:         "TimeOnly",
	TimeFormat_Custom:           "Custom",
	TimeFormat_CustomGO:         "CustomGO",
	TimeFormat_Unix_Secs:        "UnixSecs",
	TimeFormat_Unix_Milli:       "UnixMilli",
	TimeFormat_Unix_Micro:       "UnixMicro",
	TimeFormat_Unix_Nano:        "UnixNano",
}

var TimeFormatToLayout = map[TimeFormat]string{
	TimeFormat_ANSIC:            "Mon Jan _2 15:04:05 2006",
	TimeFormat_UnixDate:         "Mon Jan _2 15:04:05 MST 2006",
	TimeFormat_RubyDate:         "Mon Jan 02 15:04:05 -0700 2006",
	TimeFormat_RFC822:           "02 Jan 06 15:04 MST",
	TimeFormat_RFC822Z:          "02 Jan 06 15:04 -0700",
	TimeFormat_RFC850:           "Monday, 02-Jan-06 15:04:05 MST",
	TimeFormat_RFC1123:          "Mon, 02 Jan 2006 15:04:05 MST",
	TimeFormat_RFC1123Z:         "Mon, 02 Jan 2006 15:04:05 -0700",
	TimeFormat_RFC3339:          "2006-01-02T15:04:05Z07:00",
	TimeFormat_RFC3339Nano:      "2006-01-02T15:04:05.999999999Z07:00",
	TimeFormat_Kitchen:          "3:04PM",
	TimeFormat_Stamp:            "Jan _2 15:04:05",
	TimeFormat_StampMilli:       "Jan _2 15:04:05.000",
	TimeFormat_StampMicro:       "Jan _2 15:04:05.000000",
	TimeFormat_StampNano:        "Jan _2 15:04:05.000000000",
	TimeFormat_USDateTime:       "2006-01-02 15:04:05",
	TimeFormat_USDateTimeZ:      "2006-01-02 15:04:05 -0700",
	TimeFormat_USDateTimeMilliZ: "2006-01-02 15:04:05.000 -0700",
	TimeFormat_USDateTimeMicroZ: "2006-01-02 15:04:05.000000 -0700",
	TimeFormat_USDateTimeNanoZ:  "2006-01-02 15:04:05.000000000 -0700",
	TimeFormat_USDateShort:      "1/2/06",
	TimeFormat_USDate:           "01/02/2006",
	TimeFormat_EUDateTime:       "2006-02-01 15:04:05",
	TimeFormat_EUDateTimeZ:      "2006-02-01 15:04:05 -0700",
	TimeFormat_EUDateTimeMilliZ: "2006-02-01 15:04:05.000 -0700",
	TimeFormat_EUDateTimeMicroZ: "2006-02-01 15:04:05.000000 -0700",
	TimeFormat_EUDateTimeNanoZ:  "2006-02-01 15:04:05.000000000 -0700",
	TimeFormat_EUDateShort:      "2/1/06",
	TimeFormat_EUDate:           "02/01/2006",
	TimeFormat_DateOnly:         "2006-01-02",
	TimeFormat_TimeOnly:         "15:04:05",
}

type tokenFragment struct {
	text        string
	isSeparator bool
}

var EntityToGoMap = map[string]string{
	"yy":        "06",
	"yyyy":      "2006",
	"m":         "1",
	"mm":        "01",
	"mmm":       "Jan",
	"mmmm":      "January",
	"d":         "2",
	"dd":        "02",
	"ddd":       "Mon",
	"dddd":      "Monday",
	"h":         "3",
	"hh":        "03",
	"hhh":       "15",
	"n":         "4",
	"nn":        "04",
	"s":         "5",
	"ss":        "05",
	"zzz":       "000",
	"zzzzzz":    "000000",
	"zzzzzzzzz": "000000000",

	// yes, these are redundant, but...
	// I wanted to formalize these time encodings as entities, so users don't have to specify as separators
	"am":         "PM",
	"pm":         "PM",        // go allows using am or pm for the same thing
	"000":        "000",       // milliseconds
	"000000":     "000000",    // Microseconds
	"000000000":  "000000000", // Nanoseconds
	"thh":        "-07",
	"thhmm":      "-0700",
	"thh%mm":     "-07:00",
	"thhmmss":    "-070000",
	"thh%mm%ss":  "-07:00:00",
	"zthh":       "Z-07",
	"zthhmm":     "Z-0700",
	"zthh%mm":    "Z-07:00",
	"zthhmmss":   "Z-070000",
	"zthh%mm%ss": "Z-07:00:00",
}

type EntityDescription struct {
	Name        string
	Description string
}

var EntityToDescription = []EntityDescription{
	{"yy", "Two digit date"},
	{"yyyy", "Four digit date"},
	{"m", "Single digit month.  Shows as 2 digits for values over 9."},
	{"mm", "Double digit month"},
	{"mmm", "Three char month abbreviation: Jan, Feb, etc"},
	{"mmmm", "Full month name"},
	{"d", "Single digit day of month.  Shows as two digits for values over 9."},
	{"dd", "Double digit day of month"},
	{"ddd", "Three letter day abbreviation, Mon, Tue, etc"},
	{"dddd", "Full day name"},
	{"h", "Single digit hour value for 12 hour clock.  Shows as 2 digits for values over 9"},
	{"hh", "Double digit hour value for 12 hour clock."},
	{"hhh", "Double digit hour value for 24 hour clock."},
	{"n", "Single digit minute.  Shows as two digits for values over 9."},
	{"nn", "Double digit minute."},
	{"s", "Single digit seconds.  Shows as two digits for values over 9."},
	{"ss", "Double digit seconds."},
	{"zzz", "Milliseconds"},
	{"zzzzzz", "Microseconds"},
	{"zzzzzzzzz", "Nanoseconds"},
	{"am", "Shows AM or PM as appropriate."},
	{"pm", "Shows AM or PM as appropriate."},
	{"000", "Milliseconds"},
	{"000000", "Microseconds"},
	{"000000000", "Nanoseconds"},
	{"thh", "Timezone with hours only, e.g. -07"},
	{"thhmm", "Timezone with hours and mins, e.g. -0700"},
	{"thh%mm", "Timezone with hours and mins using separator, e.g. -07,00"},
	{"thhmmss", "Timezone with hours, mins, and seconds, without a separator, e.g. -070000"},
	{"thh%mm%ss", "Timezone with hours, mins, and seconds, using separator, e.g. -07,00,00"},
	{"zthh", "ISO 8601 Timezone using Z formatting, hours only. E.G. Z-07"},
	{"zthhmm", "ISO 8601 Timezone using Z formatting, hours and mins, e.g. Z-0700"},
	{"zthh%mm", "ISO 8601 Timezone using Z formatting, hours and mins with separator, e.q. Z-07,00"},
	{"zthhmmss", "ISO 8601 Timezone using Z formatting, hours, mins and seconds, e.g. Z-070000"},
	{"zthh%mm%ss", "ISO 8601 Timezone using Z formatting, hours, mins and seconds with separator, e.g. Z-07,00,00"},
}
