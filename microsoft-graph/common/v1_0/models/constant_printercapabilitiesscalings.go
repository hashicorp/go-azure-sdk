package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesScalings string

const (
	PrinterCapabilitiesScalingsauto        PrinterCapabilitiesScalings = "Auto"
	PrinterCapabilitiesScalingsfill        PrinterCapabilitiesScalings = "Fill"
	PrinterCapabilitiesScalingsfit         PrinterCapabilitiesScalings = "Fit"
	PrinterCapabilitiesScalingsnone        PrinterCapabilitiesScalings = "None"
	PrinterCapabilitiesScalingsshrinkToFit PrinterCapabilitiesScalings = "ShrinkToFit"
)

func PossibleValuesForPrinterCapabilitiesScalings() []string {
	return []string{
		string(PrinterCapabilitiesScalingsauto),
		string(PrinterCapabilitiesScalingsfill),
		string(PrinterCapabilitiesScalingsfit),
		string(PrinterCapabilitiesScalingsnone),
		string(PrinterCapabilitiesScalingsshrinkToFit),
	}
}

func parsePrinterCapabilitiesScalings(input string) (*PrinterCapabilitiesScalings, error) {
	vals := map[string]PrinterCapabilitiesScalings{
		"auto":        PrinterCapabilitiesScalingsauto,
		"fill":        PrinterCapabilitiesScalingsfill,
		"fit":         PrinterCapabilitiesScalingsfit,
		"none":        PrinterCapabilitiesScalingsnone,
		"shrinktofit": PrinterCapabilitiesScalingsshrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesScalings(input)
	return &out, nil
}
