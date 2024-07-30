package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsDuplexMode string

const (
	PrinterDefaultsDuplexMode_FlipOnLongEdge  PrinterDefaultsDuplexMode = "flipOnLongEdge"
	PrinterDefaultsDuplexMode_FlipOnShortEdge PrinterDefaultsDuplexMode = "flipOnShortEdge"
	PrinterDefaultsDuplexMode_OneSided        PrinterDefaultsDuplexMode = "oneSided"
)

func PossibleValuesForPrinterDefaultsDuplexMode() []string {
	return []string{
		string(PrinterDefaultsDuplexMode_FlipOnLongEdge),
		string(PrinterDefaultsDuplexMode_FlipOnShortEdge),
		string(PrinterDefaultsDuplexMode_OneSided),
	}
}

func (s *PrinterDefaultsDuplexMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsDuplexMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsDuplexMode(input string) (*PrinterDefaultsDuplexMode, error) {
	vals := map[string]PrinterDefaultsDuplexMode{
		"fliponlongedge":  PrinterDefaultsDuplexMode_FlipOnLongEdge,
		"fliponshortedge": PrinterDefaultsDuplexMode_FlipOnShortEdge,
		"onesided":        PrinterDefaultsDuplexMode_OneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsDuplexMode(input)
	return &out, nil
}
