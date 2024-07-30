package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationDuplexMode string

const (
	PrinterDocumentConfigurationDuplexMode_FlipOnLongEdge  PrinterDocumentConfigurationDuplexMode = "flipOnLongEdge"
	PrinterDocumentConfigurationDuplexMode_FlipOnShortEdge PrinterDocumentConfigurationDuplexMode = "flipOnShortEdge"
	PrinterDocumentConfigurationDuplexMode_OneSided        PrinterDocumentConfigurationDuplexMode = "oneSided"
)

func PossibleValuesForPrinterDocumentConfigurationDuplexMode() []string {
	return []string{
		string(PrinterDocumentConfigurationDuplexMode_FlipOnLongEdge),
		string(PrinterDocumentConfigurationDuplexMode_FlipOnShortEdge),
		string(PrinterDocumentConfigurationDuplexMode_OneSided),
	}
}

func (s *PrinterDocumentConfigurationDuplexMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationDuplexMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationDuplexMode(input string) (*PrinterDocumentConfigurationDuplexMode, error) {
	vals := map[string]PrinterDocumentConfigurationDuplexMode{
		"fliponlongedge":  PrinterDocumentConfigurationDuplexMode_FlipOnLongEdge,
		"fliponshortedge": PrinterDocumentConfigurationDuplexMode_FlipOnShortEdge,
		"onesided":        PrinterDocumentConfigurationDuplexMode_OneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationDuplexMode(input)
	return &out, nil
}
