package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsQuality string

const (
	PrinterDefaultsQuality_High   PrinterDefaultsQuality = "high"
	PrinterDefaultsQuality_Low    PrinterDefaultsQuality = "low"
	PrinterDefaultsQuality_Medium PrinterDefaultsQuality = "medium"
)

func PossibleValuesForPrinterDefaultsQuality() []string {
	return []string{
		string(PrinterDefaultsQuality_High),
		string(PrinterDefaultsQuality_Low),
		string(PrinterDefaultsQuality_Medium),
	}
}

func (s *PrinterDefaultsQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsQuality(input string) (*PrinterDefaultsQuality, error) {
	vals := map[string]PrinterDefaultsQuality{
		"high":   PrinterDefaultsQuality_High,
		"low":    PrinterDefaultsQuality_Low,
		"medium": PrinterDefaultsQuality_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsQuality(input)
	return &out, nil
}
