package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesMultipageLayouts string

const (
	PrinterCapabilitiesMultipageLayouts_ClockwiseFromBottomLeft         PrinterCapabilitiesMultipageLayouts = "clockwiseFromBottomLeft"
	PrinterCapabilitiesMultipageLayouts_ClockwiseFromBottomRight        PrinterCapabilitiesMultipageLayouts = "clockwiseFromBottomRight"
	PrinterCapabilitiesMultipageLayouts_ClockwiseFromTopLeft            PrinterCapabilitiesMultipageLayouts = "clockwiseFromTopLeft"
	PrinterCapabilitiesMultipageLayouts_ClockwiseFromTopRight           PrinterCapabilitiesMultipageLayouts = "clockwiseFromTopRight"
	PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromBottomLeft  PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromBottomLeft"
	PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromBottomRight PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromBottomRight"
	PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromTopLeft     PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromTopLeft"
	PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromTopRight    PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrinterCapabilitiesMultipageLayouts() []string {
	return []string{
		string(PrinterCapabilitiesMultipageLayouts_ClockwiseFromBottomLeft),
		string(PrinterCapabilitiesMultipageLayouts_ClockwiseFromBottomRight),
		string(PrinterCapabilitiesMultipageLayouts_ClockwiseFromTopLeft),
		string(PrinterCapabilitiesMultipageLayouts_ClockwiseFromTopRight),
		string(PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromBottomLeft),
		string(PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromBottomRight),
		string(PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromTopLeft),
		string(PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromTopRight),
	}
}

func (s *PrinterCapabilitiesMultipageLayouts) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesMultipageLayouts(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesMultipageLayouts(input string) (*PrinterCapabilitiesMultipageLayouts, error) {
	vals := map[string]PrinterCapabilitiesMultipageLayouts{
		"clockwisefrombottomleft":         PrinterCapabilitiesMultipageLayouts_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterCapabilitiesMultipageLayouts_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterCapabilitiesMultipageLayouts_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterCapabilitiesMultipageLayouts_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterCapabilitiesMultipageLayouts_CounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesMultipageLayouts(input)
	return &out, nil
}
