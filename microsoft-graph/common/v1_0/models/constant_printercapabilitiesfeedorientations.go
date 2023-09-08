package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesFeedOrientations string

const (
	PrinterCapabilitiesFeedOrientationslongEdgeFirst  PrinterCapabilitiesFeedOrientations = "LongEdgeFirst"
	PrinterCapabilitiesFeedOrientationsshortEdgeFirst PrinterCapabilitiesFeedOrientations = "ShortEdgeFirst"
)

func PossibleValuesForPrinterCapabilitiesFeedOrientations() []string {
	return []string{
		string(PrinterCapabilitiesFeedOrientationslongEdgeFirst),
		string(PrinterCapabilitiesFeedOrientationsshortEdgeFirst),
	}
}

func parsePrinterCapabilitiesFeedOrientations(input string) (*PrinterCapabilitiesFeedOrientations, error) {
	vals := map[string]PrinterCapabilitiesFeedOrientations{
		"longedgefirst":  PrinterCapabilitiesFeedOrientationslongEdgeFirst,
		"shortedgefirst": PrinterCapabilitiesFeedOrientationsshortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFeedOrientations(input)
	return &out, nil
}
