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

type DateTimeFormatter struct {
	dateTime time.Time
}

func NewDateTimeFormatter(dateTime time.Time) *DateTimeFormatter {
	return &DateTimeFormatter{dateTime: dateTime}
}

func (dtf *DateTimeFormatter) FormatDateTime(outputFormat TimeFormat) (formattedDate string, err error) {
	layout := "USDateTimeZ"
	found := false

	switch outputFormat {
	case TimeFormat_Unix_Secs:
		unixInt := dtf.dateTime.Unix()
		return strconv.FormatInt(unixInt, 10), nil
	case TimeFormat_Unix_Milli:
		unixInt := dtf.dateTime.UnixMilli()
		return strconv.FormatInt(unixInt, 10), nil
	case TimeFormat_Unix_Micro:
		unixInt := dtf.dateTime.UnixMicro()
		return strconv.FormatInt(unixInt, 10), nil
	case TimeFormat_Unix_Nano:
		unixInt := dtf.dateTime.UnixNano()
		return strconv.FormatInt(unixInt, 10), nil
	case TimeFormat_Custom:
		layout, err = dtf.BuildCustomLayout(CmdHelpers.OutputLayout)
		if err != nil {
			return "", err
		}
	case TimeFormat_CustomGO:
		layout = CmdHelpers.OutputLayout
	default:
		layout, found = TimeFormatToLayout[outputFormat]
		if !found {
			return "", fmt.Errorf("Unknown output format reference: %d", int(outputFormat))
		}
	}

	return dtf.dateTime.Format(layout), nil
}

// BuildCustomLayout receives a layout pattern in an alpha syntax and returns a layout text in a go syntax
func (dtf *DateTimeFormatter) BuildCustomLayout(layoutText string) (convertedlayout string, err error) {
	if layoutText == "" {
		return "", nil
	}

	// First, we tokenize the layout by text fragments
	currentChar := 0
	separatorMode := false
	separatorText := ""
	tokenText := ""
	var tokens []tokenFragment
	for {
		if !strings.ContainsAny(strings.ToLower(string(layoutText[currentChar])), "pymdhnszt%") {
			if !separatorMode {
				if tokenText != "" {
					// We have started a break in the layout text and there was a token construction being formed.
					// So we save it.
					tokens = append(tokens, tokenFragment{tokenText, false})
					tokenText = ""
				}

				separatorMode = true
				separatorText = string(layoutText[currentChar])
			} else {
				separatorText += string(layoutText[currentChar])
			}

			currentChar += 1
			if currentChar >= len(layoutText) {
				// we know that the separator text has at least one char in it...
				// so no need to check if the separator text has any data in it
				tokens = append(tokens, tokenFragment{separatorText, true})
				break
			}

			continue
		}

		if separatorMode {
			// really, you could never be in separatorMode and have an empty separatorText...
			// but we check for it just to be safe
			if separatorText != "" {
				tokens = append(tokens, tokenFragment{separatorText, true})
				separatorText = ""
			}
			separatorMode = false
		}

		if tokenText == "" {
			tokenText = string(layoutText[currentChar])
		} else {
			tokenText += string(layoutText[currentChar])
		}

		currentChar += 1
		if currentChar >= len(layoutText) {
			tokens = append(tokens, tokenFragment{tokenText, false})
			break
		}
	}

	// now, we re-assemble the tokens, replacing non-separator entities as we go
	goText := ""
	known := false
	for _, entity := range tokens {
		if entity.isSeparator {
			convertedlayout += entity.text
			continue
		}

		goText, known = EntityToGoMap[strings.ToLower(entity.text)]
		if !known {
			return "", fmt.Errorf("Unknown entity in custom text: %s", entity.text)
		}

		convertedlayout += goText
	}

	return convertedlayout, nil
}

func IsUnixTimeFormat(format TimeFormat) bool {
	unixTypes := []TimeFormat{
		TimeFormat_Unix_Secs,
		TimeFormat_Unix_Milli,
		TimeFormat_Unix_Micro,
		TimeFormat_Unix_Nano,
	}

	for _, formatType := range unixTypes {
		if format == formatType {
			return true
		}
	}

	return false
}
