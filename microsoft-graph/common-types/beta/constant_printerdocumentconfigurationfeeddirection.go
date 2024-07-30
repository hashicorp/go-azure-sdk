package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationFeedDirection string

const (
	PrinterDocumentConfigurationFeedDirection_LongEdgeFirst  PrinterDocumentConfigurationFeedDirection = "longEdgeFirst"
	PrinterDocumentConfigurationFeedDirection_ShortEdgeFirst PrinterDocumentConfigurationFeedDirection = "shortEdgeFirst"
)

func PossibleValuesForPrinterDocumentConfigurationFeedDirection() []string {
	return []string{
		string(PrinterDocumentConfigurationFeedDirection_LongEdgeFirst),
		string(PrinterDocumentConfigurationFeedDirection_ShortEdgeFirst),
	}
}

func (s *PrinterDocumentConfigurationFeedDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationFeedDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationFeedDirection(input string) (*PrinterDocumentConfigurationFeedDirection, error) {
	vals := map[string]PrinterDocumentConfigurationFeedDirection{
		"longedgefirst":  PrinterDocumentConfigurationFeedDirection_LongEdgeFirst,
		"shortedgefirst": PrinterDocumentConfigurationFeedDirection_ShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFeedDirection(input)
	return &out, nil
}
