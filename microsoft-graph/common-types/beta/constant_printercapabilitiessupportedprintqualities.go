package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedPrintQualities string

const (
	PrinterCapabilitiesSupportedPrintQualities_High   PrinterCapabilitiesSupportedPrintQualities = "high"
	PrinterCapabilitiesSupportedPrintQualities_Low    PrinterCapabilitiesSupportedPrintQualities = "low"
	PrinterCapabilitiesSupportedPrintQualities_Medium PrinterCapabilitiesSupportedPrintQualities = "medium"
)

func PossibleValuesForPrinterCapabilitiesSupportedPrintQualities() []string {
	return []string{
		string(PrinterCapabilitiesSupportedPrintQualities_High),
		string(PrinterCapabilitiesSupportedPrintQualities_Low),
		string(PrinterCapabilitiesSupportedPrintQualities_Medium),
	}
}

func (s *PrinterCapabilitiesSupportedPrintQualities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedPrintQualities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedPrintQualities(input string) (*PrinterCapabilitiesSupportedPrintQualities, error) {
	vals := map[string]PrinterCapabilitiesSupportedPrintQualities{
		"high":   PrinterCapabilitiesSupportedPrintQualities_High,
		"low":    PrinterCapabilitiesSupportedPrintQualities_Low,
		"medium": PrinterCapabilitiesSupportedPrintQualities_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedPrintQualities(input)
	return &out, nil
}
