package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesColorModes string

const (
	PrinterCapabilitiesColorModesauto          PrinterCapabilitiesColorModes = "Auto"
	PrinterCapabilitiesColorModesblackAndWhite PrinterCapabilitiesColorModes = "BlackAndWhite"
	PrinterCapabilitiesColorModescolor         PrinterCapabilitiesColorModes = "Color"
	PrinterCapabilitiesColorModesgrayscale     PrinterCapabilitiesColorModes = "Grayscale"
)

func PossibleValuesForPrinterCapabilitiesColorModes() []string {
	return []string{
		string(PrinterCapabilitiesColorModesauto),
		string(PrinterCapabilitiesColorModesblackAndWhite),
		string(PrinterCapabilitiesColorModescolor),
		string(PrinterCapabilitiesColorModesgrayscale),
	}
}

func parsePrinterCapabilitiesColorModes(input string) (*PrinterCapabilitiesColorModes, error) {
	vals := map[string]PrinterCapabilitiesColorModes{
		"auto":          PrinterCapabilitiesColorModesauto,
		"blackandwhite": PrinterCapabilitiesColorModesblackAndWhite,
		"color":         PrinterCapabilitiesColorModescolor,
		"grayscale":     PrinterCapabilitiesColorModesgrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesColorModes(input)
	return &out, nil
}
