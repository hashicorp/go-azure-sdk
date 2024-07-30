package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationFeedOrientation string

const (
	PrinterDocumentConfigurationFeedOrientation_LongEdgeFirst  PrinterDocumentConfigurationFeedOrientation = "longEdgeFirst"
	PrinterDocumentConfigurationFeedOrientation_ShortEdgeFirst PrinterDocumentConfigurationFeedOrientation = "shortEdgeFirst"
)

func PossibleValuesForPrinterDocumentConfigurationFeedOrientation() []string {
	return []string{
		string(PrinterDocumentConfigurationFeedOrientation_LongEdgeFirst),
		string(PrinterDocumentConfigurationFeedOrientation_ShortEdgeFirst),
	}
}

func (s *PrinterDocumentConfigurationFeedOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationFeedOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationFeedOrientation(input string) (*PrinterDocumentConfigurationFeedOrientation, error) {
	vals := map[string]PrinterDocumentConfigurationFeedOrientation{
		"longedgefirst":  PrinterDocumentConfigurationFeedOrientation_LongEdgeFirst,
		"shortedgefirst": PrinterDocumentConfigurationFeedOrientation_ShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFeedOrientation(input)
	return &out, nil
}
