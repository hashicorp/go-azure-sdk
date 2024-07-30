package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationMultipageLayout string

const (
	PrinterDocumentConfigurationMultipageLayout_ClockwiseFromBottomLeft         PrinterDocumentConfigurationMultipageLayout = "clockwiseFromBottomLeft"
	PrinterDocumentConfigurationMultipageLayout_ClockwiseFromBottomRight        PrinterDocumentConfigurationMultipageLayout = "clockwiseFromBottomRight"
	PrinterDocumentConfigurationMultipageLayout_ClockwiseFromTopLeft            PrinterDocumentConfigurationMultipageLayout = "clockwiseFromTopLeft"
	PrinterDocumentConfigurationMultipageLayout_ClockwiseFromTopRight           PrinterDocumentConfigurationMultipageLayout = "clockwiseFromTopRight"
	PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromBottomLeft  PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromBottomLeft"
	PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromBottomRight PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromBottomRight"
	PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromTopLeft     PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromTopLeft"
	PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromTopRight    PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrinterDocumentConfigurationMultipageLayout() []string {
	return []string{
		string(PrinterDocumentConfigurationMultipageLayout_ClockwiseFromBottomLeft),
		string(PrinterDocumentConfigurationMultipageLayout_ClockwiseFromBottomRight),
		string(PrinterDocumentConfigurationMultipageLayout_ClockwiseFromTopLeft),
		string(PrinterDocumentConfigurationMultipageLayout_ClockwiseFromTopRight),
		string(PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromBottomLeft),
		string(PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromBottomRight),
		string(PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromTopLeft),
		string(PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromTopRight),
	}
}

func (s *PrinterDocumentConfigurationMultipageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationMultipageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationMultipageLayout(input string) (*PrinterDocumentConfigurationMultipageLayout, error) {
	vals := map[string]PrinterDocumentConfigurationMultipageLayout{
		"clockwisefrombottomleft":         PrinterDocumentConfigurationMultipageLayout_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDocumentConfigurationMultipageLayout_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDocumentConfigurationMultipageLayout_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDocumentConfigurationMultipageLayout_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDocumentConfigurationMultipageLayout_CounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationMultipageLayout(input)
	return &out, nil
}
