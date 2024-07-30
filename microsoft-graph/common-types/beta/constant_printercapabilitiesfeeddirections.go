package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesFeedDirections string

const (
	PrinterCapabilitiesFeedDirections_LongEdgeFirst  PrinterCapabilitiesFeedDirections = "longEdgeFirst"
	PrinterCapabilitiesFeedDirections_ShortEdgeFirst PrinterCapabilitiesFeedDirections = "shortEdgeFirst"
)

func PossibleValuesForPrinterCapabilitiesFeedDirections() []string {
	return []string{
		string(PrinterCapabilitiesFeedDirections_LongEdgeFirst),
		string(PrinterCapabilitiesFeedDirections_ShortEdgeFirst),
	}
}

func (s *PrinterCapabilitiesFeedDirections) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesFeedDirections(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesFeedDirections(input string) (*PrinterCapabilitiesFeedDirections, error) {
	vals := map[string]PrinterCapabilitiesFeedDirections{
		"longedgefirst":  PrinterCapabilitiesFeedDirections_LongEdgeFirst,
		"shortedgefirst": PrinterCapabilitiesFeedDirections_ShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFeedDirections(input)
	return &out, nil
}
