package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationMultipageLayout string

const (
	PrintJobConfigurationMultipageLayoutclockwiseFromBottomLeft         PrintJobConfigurationMultipageLayout = "ClockwiseFromBottomLeft"
	PrintJobConfigurationMultipageLayoutclockwiseFromBottomRight        PrintJobConfigurationMultipageLayout = "ClockwiseFromBottomRight"
	PrintJobConfigurationMultipageLayoutclockwiseFromTopLeft            PrintJobConfigurationMultipageLayout = "ClockwiseFromTopLeft"
	PrintJobConfigurationMultipageLayoutclockwiseFromTopRight           PrintJobConfigurationMultipageLayout = "ClockwiseFromTopRight"
	PrintJobConfigurationMultipageLayoutcounterclockwiseFromBottomLeft  PrintJobConfigurationMultipageLayout = "CounterclockwiseFromBottomLeft"
	PrintJobConfigurationMultipageLayoutcounterclockwiseFromBottomRight PrintJobConfigurationMultipageLayout = "CounterclockwiseFromBottomRight"
	PrintJobConfigurationMultipageLayoutcounterclockwiseFromTopLeft     PrintJobConfigurationMultipageLayout = "CounterclockwiseFromTopLeft"
	PrintJobConfigurationMultipageLayoutcounterclockwiseFromTopRight    PrintJobConfigurationMultipageLayout = "CounterclockwiseFromTopRight"
)

func PossibleValuesForPrintJobConfigurationMultipageLayout() []string {
	return []string{
		string(PrintJobConfigurationMultipageLayoutclockwiseFromBottomLeft),
		string(PrintJobConfigurationMultipageLayoutclockwiseFromBottomRight),
		string(PrintJobConfigurationMultipageLayoutclockwiseFromTopLeft),
		string(PrintJobConfigurationMultipageLayoutclockwiseFromTopRight),
		string(PrintJobConfigurationMultipageLayoutcounterclockwiseFromBottomLeft),
		string(PrintJobConfigurationMultipageLayoutcounterclockwiseFromBottomRight),
		string(PrintJobConfigurationMultipageLayoutcounterclockwiseFromTopLeft),
		string(PrintJobConfigurationMultipageLayoutcounterclockwiseFromTopRight),
	}
}

func parsePrintJobConfigurationMultipageLayout(input string) (*PrintJobConfigurationMultipageLayout, error) {
	vals := map[string]PrintJobConfigurationMultipageLayout{
		"clockwisefrombottomleft":         PrintJobConfigurationMultipageLayoutclockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrintJobConfigurationMultipageLayoutclockwiseFromBottomRight,
		"clockwisefromtopleft":            PrintJobConfigurationMultipageLayoutclockwiseFromTopLeft,
		"clockwisefromtopright":           PrintJobConfigurationMultipageLayoutclockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrintJobConfigurationMultipageLayoutcounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrintJobConfigurationMultipageLayoutcounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrintJobConfigurationMultipageLayoutcounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrintJobConfigurationMultipageLayoutcounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationMultipageLayout(input)
	return &out, nil
}
