package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationMultipageLayout string

const (
	PrintJobConfigurationMultipageLayout_ClockwiseFromBottomLeft         PrintJobConfigurationMultipageLayout = "clockwiseFromBottomLeft"
	PrintJobConfigurationMultipageLayout_ClockwiseFromBottomRight        PrintJobConfigurationMultipageLayout = "clockwiseFromBottomRight"
	PrintJobConfigurationMultipageLayout_ClockwiseFromTopLeft            PrintJobConfigurationMultipageLayout = "clockwiseFromTopLeft"
	PrintJobConfigurationMultipageLayout_ClockwiseFromTopRight           PrintJobConfigurationMultipageLayout = "clockwiseFromTopRight"
	PrintJobConfigurationMultipageLayout_CounterclockwiseFromBottomLeft  PrintJobConfigurationMultipageLayout = "counterclockwiseFromBottomLeft"
	PrintJobConfigurationMultipageLayout_CounterclockwiseFromBottomRight PrintJobConfigurationMultipageLayout = "counterclockwiseFromBottomRight"
	PrintJobConfigurationMultipageLayout_CounterclockwiseFromTopLeft     PrintJobConfigurationMultipageLayout = "counterclockwiseFromTopLeft"
	PrintJobConfigurationMultipageLayout_CounterclockwiseFromTopRight    PrintJobConfigurationMultipageLayout = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrintJobConfigurationMultipageLayout() []string {
	return []string{
		string(PrintJobConfigurationMultipageLayout_ClockwiseFromBottomLeft),
		string(PrintJobConfigurationMultipageLayout_ClockwiseFromBottomRight),
		string(PrintJobConfigurationMultipageLayout_ClockwiseFromTopLeft),
		string(PrintJobConfigurationMultipageLayout_ClockwiseFromTopRight),
		string(PrintJobConfigurationMultipageLayout_CounterclockwiseFromBottomLeft),
		string(PrintJobConfigurationMultipageLayout_CounterclockwiseFromBottomRight),
		string(PrintJobConfigurationMultipageLayout_CounterclockwiseFromTopLeft),
		string(PrintJobConfigurationMultipageLayout_CounterclockwiseFromTopRight),
	}
}

func (s *PrintJobConfigurationMultipageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationMultipageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationMultipageLayout(input string) (*PrintJobConfigurationMultipageLayout, error) {
	vals := map[string]PrintJobConfigurationMultipageLayout{
		"clockwisefrombottomleft":         PrintJobConfigurationMultipageLayout_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrintJobConfigurationMultipageLayout_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrintJobConfigurationMultipageLayout_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrintJobConfigurationMultipageLayout_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrintJobConfigurationMultipageLayout_CounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrintJobConfigurationMultipageLayout_CounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrintJobConfigurationMultipageLayout_CounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrintJobConfigurationMultipageLayout_CounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationMultipageLayout(input)
	return &out, nil
}
