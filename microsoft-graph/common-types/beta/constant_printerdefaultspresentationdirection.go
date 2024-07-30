package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsPresentationDirection string

const (
	PrinterDefaultsPresentationDirection_ClockwiseFromBottomLeft         PrinterDefaultsPresentationDirection = "clockwiseFromBottomLeft"
	PrinterDefaultsPresentationDirection_ClockwiseFromBottomRight        PrinterDefaultsPresentationDirection = "clockwiseFromBottomRight"
	PrinterDefaultsPresentationDirection_ClockwiseFromTopLeft            PrinterDefaultsPresentationDirection = "clockwiseFromTopLeft"
	PrinterDefaultsPresentationDirection_ClockwiseFromTopRight           PrinterDefaultsPresentationDirection = "clockwiseFromTopRight"
	PrinterDefaultsPresentationDirection_CounterClockwiseFromBottomLeft  PrinterDefaultsPresentationDirection = "counterClockwiseFromBottomLeft"
	PrinterDefaultsPresentationDirection_CounterClockwiseFromBottomRight PrinterDefaultsPresentationDirection = "counterClockwiseFromBottomRight"
	PrinterDefaultsPresentationDirection_CounterClockwiseFromTopLeft     PrinterDefaultsPresentationDirection = "counterClockwiseFromTopLeft"
	PrinterDefaultsPresentationDirection_CounterClockwiseFromTopRight    PrinterDefaultsPresentationDirection = "counterClockwiseFromTopRight"
)

func PossibleValuesForPrinterDefaultsPresentationDirection() []string {
	return []string{
		string(PrinterDefaultsPresentationDirection_ClockwiseFromBottomLeft),
		string(PrinterDefaultsPresentationDirection_ClockwiseFromBottomRight),
		string(PrinterDefaultsPresentationDirection_ClockwiseFromTopLeft),
		string(PrinterDefaultsPresentationDirection_ClockwiseFromTopRight),
		string(PrinterDefaultsPresentationDirection_CounterClockwiseFromBottomLeft),
		string(PrinterDefaultsPresentationDirection_CounterClockwiseFromBottomRight),
		string(PrinterDefaultsPresentationDirection_CounterClockwiseFromTopLeft),
		string(PrinterDefaultsPresentationDirection_CounterClockwiseFromTopRight),
	}
}

func (s *PrinterDefaultsPresentationDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsPresentationDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsPresentationDirection(input string) (*PrinterDefaultsPresentationDirection, error) {
	vals := map[string]PrinterDefaultsPresentationDirection{
		"clockwisefrombottomleft":         PrinterDefaultsPresentationDirection_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDefaultsPresentationDirection_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDefaultsPresentationDirection_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDefaultsPresentationDirection_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDefaultsPresentationDirection_CounterClockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDefaultsPresentationDirection_CounterClockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDefaultsPresentationDirection_CounterClockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDefaultsPresentationDirection_CounterClockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPresentationDirection(input)
	return &out, nil
}
