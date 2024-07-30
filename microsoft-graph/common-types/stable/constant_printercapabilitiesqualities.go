package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesQualities string

const (
	PrinterCapabilitiesQualities_High   PrinterCapabilitiesQualities = "high"
	PrinterCapabilitiesQualities_Low    PrinterCapabilitiesQualities = "low"
	PrinterCapabilitiesQualities_Medium PrinterCapabilitiesQualities = "medium"
)

func PossibleValuesForPrinterCapabilitiesQualities() []string {
	return []string{
		string(PrinterCapabilitiesQualities_High),
		string(PrinterCapabilitiesQualities_Low),
		string(PrinterCapabilitiesQualities_Medium),
	}
}

func (s *PrinterCapabilitiesQualities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesQualities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesQualities(input string) (*PrinterCapabilitiesQualities, error) {
	vals := map[string]PrinterCapabilitiesQualities{
		"high":   PrinterCapabilitiesQualities_High,
		"low":    PrinterCapabilitiesQualities_Low,
		"medium": PrinterCapabilitiesQualities_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesQualities(input)
	return &out, nil
}
