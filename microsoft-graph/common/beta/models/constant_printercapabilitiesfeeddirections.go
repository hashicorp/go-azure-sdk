package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesFeedDirections string

const (
	PrinterCapabilitiesFeedDirectionslongEdgeFirst  PrinterCapabilitiesFeedDirections = "LongEdgeFirst"
	PrinterCapabilitiesFeedDirectionsshortEdgeFirst PrinterCapabilitiesFeedDirections = "ShortEdgeFirst"
)

func PossibleValuesForPrinterCapabilitiesFeedDirections() []string {
	return []string{
		string(PrinterCapabilitiesFeedDirectionslongEdgeFirst),
		string(PrinterCapabilitiesFeedDirectionsshortEdgeFirst),
	}
}

func parsePrinterCapabilitiesFeedDirections(input string) (*PrinterCapabilitiesFeedDirections, error) {
	vals := map[string]PrinterCapabilitiesFeedDirections{
		"longedgefirst":  PrinterCapabilitiesFeedDirectionslongEdgeFirst,
		"shortedgefirst": PrinterCapabilitiesFeedDirectionsshortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFeedDirections(input)
	return &out, nil
}
