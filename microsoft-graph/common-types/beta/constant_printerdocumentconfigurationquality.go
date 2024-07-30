package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationQuality string

const (
	PrinterDocumentConfigurationQuality_High   PrinterDocumentConfigurationQuality = "high"
	PrinterDocumentConfigurationQuality_Low    PrinterDocumentConfigurationQuality = "low"
	PrinterDocumentConfigurationQuality_Medium PrinterDocumentConfigurationQuality = "medium"
)

func PossibleValuesForPrinterDocumentConfigurationQuality() []string {
	return []string{
		string(PrinterDocumentConfigurationQuality_High),
		string(PrinterDocumentConfigurationQuality_Low),
		string(PrinterDocumentConfigurationQuality_Medium),
	}
}

func (s *PrinterDocumentConfigurationQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationQuality(input string) (*PrinterDocumentConfigurationQuality, error) {
	vals := map[string]PrinterDocumentConfigurationQuality{
		"high":   PrinterDocumentConfigurationQuality_High,
		"low":    PrinterDocumentConfigurationQuality_Low,
		"medium": PrinterDocumentConfigurationQuality_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationQuality(input)
	return &out, nil
}
