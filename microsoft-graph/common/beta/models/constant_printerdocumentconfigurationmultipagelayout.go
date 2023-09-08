package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationMultipageLayout string

const (
	PrinterDocumentConfigurationMultipageLayoutclockwiseFromBottomLeft         PrinterDocumentConfigurationMultipageLayout = "ClockwiseFromBottomLeft"
	PrinterDocumentConfigurationMultipageLayoutclockwiseFromBottomRight        PrinterDocumentConfigurationMultipageLayout = "ClockwiseFromBottomRight"
	PrinterDocumentConfigurationMultipageLayoutclockwiseFromTopLeft            PrinterDocumentConfigurationMultipageLayout = "ClockwiseFromTopLeft"
	PrinterDocumentConfigurationMultipageLayoutclockwiseFromTopRight           PrinterDocumentConfigurationMultipageLayout = "ClockwiseFromTopRight"
	PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromBottomLeft  PrinterDocumentConfigurationMultipageLayout = "CounterclockwiseFromBottomLeft"
	PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromBottomRight PrinterDocumentConfigurationMultipageLayout = "CounterclockwiseFromBottomRight"
	PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromTopLeft     PrinterDocumentConfigurationMultipageLayout = "CounterclockwiseFromTopLeft"
	PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromTopRight    PrinterDocumentConfigurationMultipageLayout = "CounterclockwiseFromTopRight"
)

func PossibleValuesForPrinterDocumentConfigurationMultipageLayout() []string {
	return []string{
		string(PrinterDocumentConfigurationMultipageLayoutclockwiseFromBottomLeft),
		string(PrinterDocumentConfigurationMultipageLayoutclockwiseFromBottomRight),
		string(PrinterDocumentConfigurationMultipageLayoutclockwiseFromTopLeft),
		string(PrinterDocumentConfigurationMultipageLayoutclockwiseFromTopRight),
		string(PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromBottomLeft),
		string(PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromBottomRight),
		string(PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromTopLeft),
		string(PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromTopRight),
	}
}

func parsePrinterDocumentConfigurationMultipageLayout(input string) (*PrinterDocumentConfigurationMultipageLayout, error) {
	vals := map[string]PrinterDocumentConfigurationMultipageLayout{
		"clockwisefrombottomleft":         PrinterDocumentConfigurationMultipageLayoutclockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDocumentConfigurationMultipageLayoutclockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDocumentConfigurationMultipageLayoutclockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDocumentConfigurationMultipageLayoutclockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDocumentConfigurationMultipageLayoutcounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationMultipageLayout(input)
	return &out, nil
}
