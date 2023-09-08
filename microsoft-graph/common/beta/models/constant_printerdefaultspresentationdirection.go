package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsPresentationDirection string

const (
	PrinterDefaultsPresentationDirectionclockwiseFromBottomLeft         PrinterDefaultsPresentationDirection = "ClockwiseFromBottomLeft"
	PrinterDefaultsPresentationDirectionclockwiseFromBottomRight        PrinterDefaultsPresentationDirection = "ClockwiseFromBottomRight"
	PrinterDefaultsPresentationDirectionclockwiseFromTopLeft            PrinterDefaultsPresentationDirection = "ClockwiseFromTopLeft"
	PrinterDefaultsPresentationDirectionclockwiseFromTopRight           PrinterDefaultsPresentationDirection = "ClockwiseFromTopRight"
	PrinterDefaultsPresentationDirectioncounterClockwiseFromBottomLeft  PrinterDefaultsPresentationDirection = "CounterClockwiseFromBottomLeft"
	PrinterDefaultsPresentationDirectioncounterClockwiseFromBottomRight PrinterDefaultsPresentationDirection = "CounterClockwiseFromBottomRight"
	PrinterDefaultsPresentationDirectioncounterClockwiseFromTopLeft     PrinterDefaultsPresentationDirection = "CounterClockwiseFromTopLeft"
	PrinterDefaultsPresentationDirectioncounterClockwiseFromTopRight    PrinterDefaultsPresentationDirection = "CounterClockwiseFromTopRight"
)

func PossibleValuesForPrinterDefaultsPresentationDirection() []string {
	return []string{
		string(PrinterDefaultsPresentationDirectionclockwiseFromBottomLeft),
		string(PrinterDefaultsPresentationDirectionclockwiseFromBottomRight),
		string(PrinterDefaultsPresentationDirectionclockwiseFromTopLeft),
		string(PrinterDefaultsPresentationDirectionclockwiseFromTopRight),
		string(PrinterDefaultsPresentationDirectioncounterClockwiseFromBottomLeft),
		string(PrinterDefaultsPresentationDirectioncounterClockwiseFromBottomRight),
		string(PrinterDefaultsPresentationDirectioncounterClockwiseFromTopLeft),
		string(PrinterDefaultsPresentationDirectioncounterClockwiseFromTopRight),
	}
}

func parsePrinterDefaultsPresentationDirection(input string) (*PrinterDefaultsPresentationDirection, error) {
	vals := map[string]PrinterDefaultsPresentationDirection{
		"clockwisefrombottomleft":         PrinterDefaultsPresentationDirectionclockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDefaultsPresentationDirectionclockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDefaultsPresentationDirectionclockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDefaultsPresentationDirectionclockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDefaultsPresentationDirectioncounterClockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDefaultsPresentationDirectioncounterClockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDefaultsPresentationDirectioncounterClockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDefaultsPresentationDirectioncounterClockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPresentationDirection(input)
	return &out, nil
}
