package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationScaling string

const (
	PrintJobConfigurationScalingauto        PrintJobConfigurationScaling = "Auto"
	PrintJobConfigurationScalingfill        PrintJobConfigurationScaling = "Fill"
	PrintJobConfigurationScalingfit         PrintJobConfigurationScaling = "Fit"
	PrintJobConfigurationScalingnone        PrintJobConfigurationScaling = "None"
	PrintJobConfigurationScalingshrinkToFit PrintJobConfigurationScaling = "ShrinkToFit"
)

func PossibleValuesForPrintJobConfigurationScaling() []string {
	return []string{
		string(PrintJobConfigurationScalingauto),
		string(PrintJobConfigurationScalingfill),
		string(PrintJobConfigurationScalingfit),
		string(PrintJobConfigurationScalingnone),
		string(PrintJobConfigurationScalingshrinkToFit),
	}
}

func parsePrintJobConfigurationScaling(input string) (*PrintJobConfigurationScaling, error) {
	vals := map[string]PrintJobConfigurationScaling{
		"auto":        PrintJobConfigurationScalingauto,
		"fill":        PrintJobConfigurationScalingfill,
		"fit":         PrintJobConfigurationScalingfit,
		"none":        PrintJobConfigurationScalingnone,
		"shrinktofit": PrintJobConfigurationScalingshrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationScaling(input)
	return &out, nil
}
