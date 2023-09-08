package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesMultipageLayouts string

const (
	PrinterCapabilitiesMultipageLayoutsclockwiseFromBottomLeft         PrinterCapabilitiesMultipageLayouts = "ClockwiseFromBottomLeft"
	PrinterCapabilitiesMultipageLayoutsclockwiseFromBottomRight        PrinterCapabilitiesMultipageLayouts = "ClockwiseFromBottomRight"
	PrinterCapabilitiesMultipageLayoutsclockwiseFromTopLeft            PrinterCapabilitiesMultipageLayouts = "ClockwiseFromTopLeft"
	PrinterCapabilitiesMultipageLayoutsclockwiseFromTopRight           PrinterCapabilitiesMultipageLayouts = "ClockwiseFromTopRight"
	PrinterCapabilitiesMultipageLayoutscounterclockwiseFromBottomLeft  PrinterCapabilitiesMultipageLayouts = "CounterclockwiseFromBottomLeft"
	PrinterCapabilitiesMultipageLayoutscounterclockwiseFromBottomRight PrinterCapabilitiesMultipageLayouts = "CounterclockwiseFromBottomRight"
	PrinterCapabilitiesMultipageLayoutscounterclockwiseFromTopLeft     PrinterCapabilitiesMultipageLayouts = "CounterclockwiseFromTopLeft"
	PrinterCapabilitiesMultipageLayoutscounterclockwiseFromTopRight    PrinterCapabilitiesMultipageLayouts = "CounterclockwiseFromTopRight"
)

func PossibleValuesForPrinterCapabilitiesMultipageLayouts() []string {
	return []string{
		string(PrinterCapabilitiesMultipageLayoutsclockwiseFromBottomLeft),
		string(PrinterCapabilitiesMultipageLayoutsclockwiseFromBottomRight),
		string(PrinterCapabilitiesMultipageLayoutsclockwiseFromTopLeft),
		string(PrinterCapabilitiesMultipageLayoutsclockwiseFromTopRight),
		string(PrinterCapabilitiesMultipageLayoutscounterclockwiseFromBottomLeft),
		string(PrinterCapabilitiesMultipageLayoutscounterclockwiseFromBottomRight),
		string(PrinterCapabilitiesMultipageLayoutscounterclockwiseFromTopLeft),
		string(PrinterCapabilitiesMultipageLayoutscounterclockwiseFromTopRight),
	}
}

func parsePrinterCapabilitiesMultipageLayouts(input string) (*PrinterCapabilitiesMultipageLayouts, error) {
	vals := map[string]PrinterCapabilitiesMultipageLayouts{
		"clockwisefrombottomleft":         PrinterCapabilitiesMultipageLayoutsclockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterCapabilitiesMultipageLayoutsclockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterCapabilitiesMultipageLayoutsclockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterCapabilitiesMultipageLayoutsclockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterCapabilitiesMultipageLayoutscounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterCapabilitiesMultipageLayoutscounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterCapabilitiesMultipageLayoutscounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterCapabilitiesMultipageLayoutscounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesMultipageLayouts(input)
	return &out, nil
}
