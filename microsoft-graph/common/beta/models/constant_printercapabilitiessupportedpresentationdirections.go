package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedPresentationDirections string

const (
	PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromBottomLeft         PrinterCapabilitiesSupportedPresentationDirections = "ClockwiseFromBottomLeft"
	PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromBottomRight        PrinterCapabilitiesSupportedPresentationDirections = "ClockwiseFromBottomRight"
	PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromTopLeft            PrinterCapabilitiesSupportedPresentationDirections = "ClockwiseFromTopLeft"
	PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromTopRight           PrinterCapabilitiesSupportedPresentationDirections = "ClockwiseFromTopRight"
	PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromBottomLeft  PrinterCapabilitiesSupportedPresentationDirections = "CounterClockwiseFromBottomLeft"
	PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromBottomRight PrinterCapabilitiesSupportedPresentationDirections = "CounterClockwiseFromBottomRight"
	PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromTopLeft     PrinterCapabilitiesSupportedPresentationDirections = "CounterClockwiseFromTopLeft"
	PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromTopRight    PrinterCapabilitiesSupportedPresentationDirections = "CounterClockwiseFromTopRight"
)

func PossibleValuesForPrinterCapabilitiesSupportedPresentationDirections() []string {
	return []string{
		string(PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromBottomLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromBottomRight),
		string(PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromTopLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromTopRight),
		string(PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromBottomLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromBottomRight),
		string(PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromTopLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromTopRight),
	}
}

func parsePrinterCapabilitiesSupportedPresentationDirections(input string) (*PrinterCapabilitiesSupportedPresentationDirections, error) {
	vals := map[string]PrinterCapabilitiesSupportedPresentationDirections{
		"clockwisefrombottomleft":         PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterCapabilitiesSupportedPresentationDirectionsclockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterCapabilitiesSupportedPresentationDirectionscounterClockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedPresentationDirections(input)
	return &out, nil
}
