package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationScaling string

const (
	PrinterDocumentConfigurationScalingauto        PrinterDocumentConfigurationScaling = "Auto"
	PrinterDocumentConfigurationScalingfill        PrinterDocumentConfigurationScaling = "Fill"
	PrinterDocumentConfigurationScalingfit         PrinterDocumentConfigurationScaling = "Fit"
	PrinterDocumentConfigurationScalingnone        PrinterDocumentConfigurationScaling = "None"
	PrinterDocumentConfigurationScalingshrinkToFit PrinterDocumentConfigurationScaling = "ShrinkToFit"
)

func PossibleValuesForPrinterDocumentConfigurationScaling() []string {
	return []string{
		string(PrinterDocumentConfigurationScalingauto),
		string(PrinterDocumentConfigurationScalingfill),
		string(PrinterDocumentConfigurationScalingfit),
		string(PrinterDocumentConfigurationScalingnone),
		string(PrinterDocumentConfigurationScalingshrinkToFit),
	}
}

func parsePrinterDocumentConfigurationScaling(input string) (*PrinterDocumentConfigurationScaling, error) {
	vals := map[string]PrinterDocumentConfigurationScaling{
		"auto":        PrinterDocumentConfigurationScalingauto,
		"fill":        PrinterDocumentConfigurationScalingfill,
		"fit":         PrinterDocumentConfigurationScalingfit,
		"none":        PrinterDocumentConfigurationScalingnone,
		"shrinktofit": PrinterDocumentConfigurationScalingshrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationScaling(input)
	return &out, nil
}
