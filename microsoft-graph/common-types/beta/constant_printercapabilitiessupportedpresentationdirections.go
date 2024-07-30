package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedPresentationDirections string

const (
	PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromBottomLeft         PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromBottomLeft"
	PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromBottomRight        PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromBottomRight"
	PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromTopLeft            PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromTopLeft"
	PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromTopRight           PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromTopRight"
	PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromBottomLeft  PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromBottomLeft"
	PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromBottomRight PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromBottomRight"
	PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromTopLeft     PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromTopLeft"
	PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromTopRight    PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromTopRight"
)

func PossibleValuesForPrinterCapabilitiesSupportedPresentationDirections() []string {
	return []string{
		string(PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromBottomLeft),
		string(PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromBottomRight),
		string(PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromTopLeft),
		string(PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromTopRight),
		string(PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromBottomLeft),
		string(PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromBottomRight),
		string(PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromTopLeft),
		string(PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromTopRight),
	}
}

func (s *PrinterCapabilitiesSupportedPresentationDirections) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedPresentationDirections(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedPresentationDirections(input string) (*PrinterCapabilitiesSupportedPresentationDirections, error) {
	vals := map[string]PrinterCapabilitiesSupportedPresentationDirections{
		"clockwisefrombottomleft":         PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterCapabilitiesSupportedPresentationDirections_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterCapabilitiesSupportedPresentationDirections_CounterClockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedPresentationDirections(input)
	return &out, nil
}
