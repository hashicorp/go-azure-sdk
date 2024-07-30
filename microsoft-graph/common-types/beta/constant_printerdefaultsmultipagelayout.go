package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsMultipageLayout string

const (
	PrinterDefaultsMultipageLayout_ClockwiseFromBottomLeft         PrinterDefaultsMultipageLayout = "clockwiseFromBottomLeft"
	PrinterDefaultsMultipageLayout_ClockwiseFromBottomRight        PrinterDefaultsMultipageLayout = "clockwiseFromBottomRight"
	PrinterDefaultsMultipageLayout_ClockwiseFromTopLeft            PrinterDefaultsMultipageLayout = "clockwiseFromTopLeft"
	PrinterDefaultsMultipageLayout_ClockwiseFromTopRight           PrinterDefaultsMultipageLayout = "clockwiseFromTopRight"
	PrinterDefaultsMultipageLayout_CounterclockwiseFromBottomLeft  PrinterDefaultsMultipageLayout = "counterclockwiseFromBottomLeft"
	PrinterDefaultsMultipageLayout_CounterclockwiseFromBottomRight PrinterDefaultsMultipageLayout = "counterclockwiseFromBottomRight"
	PrinterDefaultsMultipageLayout_CounterclockwiseFromTopLeft     PrinterDefaultsMultipageLayout = "counterclockwiseFromTopLeft"
	PrinterDefaultsMultipageLayout_CounterclockwiseFromTopRight    PrinterDefaultsMultipageLayout = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrinterDefaultsMultipageLayout() []string {
	return []string{
		string(PrinterDefaultsMultipageLayout_ClockwiseFromBottomLeft),
		string(PrinterDefaultsMultipageLayout_ClockwiseFromBottomRight),
		string(PrinterDefaultsMultipageLayout_ClockwiseFromTopLeft),
		string(PrinterDefaultsMultipageLayout_ClockwiseFromTopRight),
		string(PrinterDefaultsMultipageLayout_CounterclockwiseFromBottomLeft),
		string(PrinterDefaultsMultipageLayout_CounterclockwiseFromBottomRight),
		string(PrinterDefaultsMultipageLayout_CounterclockwiseFromTopLeft),
		string(PrinterDefaultsMultipageLayout_CounterclockwiseFromTopRight),
	}
}

func (s *PrinterDefaultsMultipageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsMultipageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsMultipageLayout(input string) (*PrinterDefaultsMultipageLayout, error) {
	vals := map[string]PrinterDefaultsMultipageLayout{
		"clockwisefrombottomleft":         PrinterDefaultsMultipageLayout_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDefaultsMultipageLayout_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDefaultsMultipageLayout_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDefaultsMultipageLayout_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDefaultsMultipageLayout_CounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDefaultsMultipageLayout_CounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDefaultsMultipageLayout_CounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDefaultsMultipageLayout_CounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsMultipageLayout(input)
	return &out, nil
}
