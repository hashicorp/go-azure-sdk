package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsScaling string

const (
	PrinterDefaultsScalingauto        PrinterDefaultsScaling = "Auto"
	PrinterDefaultsScalingfill        PrinterDefaultsScaling = "Fill"
	PrinterDefaultsScalingfit         PrinterDefaultsScaling = "Fit"
	PrinterDefaultsScalingnone        PrinterDefaultsScaling = "None"
	PrinterDefaultsScalingshrinkToFit PrinterDefaultsScaling = "ShrinkToFit"
)

func PossibleValuesForPrinterDefaultsScaling() []string {
	return []string{
		string(PrinterDefaultsScalingauto),
		string(PrinterDefaultsScalingfill),
		string(PrinterDefaultsScalingfit),
		string(PrinterDefaultsScalingnone),
		string(PrinterDefaultsScalingshrinkToFit),
	}
}

func parsePrinterDefaultsScaling(input string) (*PrinterDefaultsScaling, error) {
	vals := map[string]PrinterDefaultsScaling{
		"auto":        PrinterDefaultsScalingauto,
		"fill":        PrinterDefaultsScalingfill,
		"fit":         PrinterDefaultsScalingfit,
		"none":        PrinterDefaultsScalingnone,
		"shrinktofit": PrinterDefaultsScalingshrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsScaling(input)
	return &out, nil
}
