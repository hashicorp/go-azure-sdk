package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsMultipageLayout string

const (
	PrinterDefaultsMultipageLayoutclockwiseFromBottomLeft         PrinterDefaultsMultipageLayout = "ClockwiseFromBottomLeft"
	PrinterDefaultsMultipageLayoutclockwiseFromBottomRight        PrinterDefaultsMultipageLayout = "ClockwiseFromBottomRight"
	PrinterDefaultsMultipageLayoutclockwiseFromTopLeft            PrinterDefaultsMultipageLayout = "ClockwiseFromTopLeft"
	PrinterDefaultsMultipageLayoutclockwiseFromTopRight           PrinterDefaultsMultipageLayout = "ClockwiseFromTopRight"
	PrinterDefaultsMultipageLayoutcounterclockwiseFromBottomLeft  PrinterDefaultsMultipageLayout = "CounterclockwiseFromBottomLeft"
	PrinterDefaultsMultipageLayoutcounterclockwiseFromBottomRight PrinterDefaultsMultipageLayout = "CounterclockwiseFromBottomRight"
	PrinterDefaultsMultipageLayoutcounterclockwiseFromTopLeft     PrinterDefaultsMultipageLayout = "CounterclockwiseFromTopLeft"
	PrinterDefaultsMultipageLayoutcounterclockwiseFromTopRight    PrinterDefaultsMultipageLayout = "CounterclockwiseFromTopRight"
)

func PossibleValuesForPrinterDefaultsMultipageLayout() []string {
	return []string{
		string(PrinterDefaultsMultipageLayoutclockwiseFromBottomLeft),
		string(PrinterDefaultsMultipageLayoutclockwiseFromBottomRight),
		string(PrinterDefaultsMultipageLayoutclockwiseFromTopLeft),
		string(PrinterDefaultsMultipageLayoutclockwiseFromTopRight),
		string(PrinterDefaultsMultipageLayoutcounterclockwiseFromBottomLeft),
		string(PrinterDefaultsMultipageLayoutcounterclockwiseFromBottomRight),
		string(PrinterDefaultsMultipageLayoutcounterclockwiseFromTopLeft),
		string(PrinterDefaultsMultipageLayoutcounterclockwiseFromTopRight),
	}
}

func parsePrinterDefaultsMultipageLayout(input string) (*PrinterDefaultsMultipageLayout, error) {
	vals := map[string]PrinterDefaultsMultipageLayout{
		"clockwisefrombottomleft":         PrinterDefaultsMultipageLayoutclockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDefaultsMultipageLayoutclockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDefaultsMultipageLayoutclockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDefaultsMultipageLayoutclockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDefaultsMultipageLayoutcounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDefaultsMultipageLayoutcounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDefaultsMultipageLayoutcounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDefaultsMultipageLayoutcounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsMultipageLayout(input)
	return &out, nil
}
