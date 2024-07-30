package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsPrintQuality string

const (
	PrinterDefaultsPrintQuality_High   PrinterDefaultsPrintQuality = "high"
	PrinterDefaultsPrintQuality_Low    PrinterDefaultsPrintQuality = "low"
	PrinterDefaultsPrintQuality_Medium PrinterDefaultsPrintQuality = "medium"
)

func PossibleValuesForPrinterDefaultsPrintQuality() []string {
	return []string{
		string(PrinterDefaultsPrintQuality_High),
		string(PrinterDefaultsPrintQuality_Low),
		string(PrinterDefaultsPrintQuality_Medium),
	}
}

func (s *PrinterDefaultsPrintQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsPrintQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsPrintQuality(input string) (*PrinterDefaultsPrintQuality, error) {
	vals := map[string]PrinterDefaultsPrintQuality{
		"high":   PrinterDefaultsPrintQuality_High,
		"low":    PrinterDefaultsPrintQuality_Low,
		"medium": PrinterDefaultsPrintQuality_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPrintQuality(input)
	return &out, nil
}
