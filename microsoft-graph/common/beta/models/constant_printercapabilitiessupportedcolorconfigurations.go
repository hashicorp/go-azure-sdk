package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedColorConfigurations string

const (
	PrinterCapabilitiesSupportedColorConfigurationsauto          PrinterCapabilitiesSupportedColorConfigurations = "Auto"
	PrinterCapabilitiesSupportedColorConfigurationsblackAndWhite PrinterCapabilitiesSupportedColorConfigurations = "BlackAndWhite"
	PrinterCapabilitiesSupportedColorConfigurationscolor         PrinterCapabilitiesSupportedColorConfigurations = "Color"
	PrinterCapabilitiesSupportedColorConfigurationsgrayscale     PrinterCapabilitiesSupportedColorConfigurations = "Grayscale"
)

func PossibleValuesForPrinterCapabilitiesSupportedColorConfigurations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedColorConfigurationsauto),
		string(PrinterCapabilitiesSupportedColorConfigurationsblackAndWhite),
		string(PrinterCapabilitiesSupportedColorConfigurationscolor),
		string(PrinterCapabilitiesSupportedColorConfigurationsgrayscale),
	}
}

func parsePrinterCapabilitiesSupportedColorConfigurations(input string) (*PrinterCapabilitiesSupportedColorConfigurations, error) {
	vals := map[string]PrinterCapabilitiesSupportedColorConfigurations{
		"auto":          PrinterCapabilitiesSupportedColorConfigurationsauto,
		"blackandwhite": PrinterCapabilitiesSupportedColorConfigurationsblackAndWhite,
		"color":         PrinterCapabilitiesSupportedColorConfigurationscolor,
		"grayscale":     PrinterCapabilitiesSupportedColorConfigurationsgrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedColorConfigurations(input)
	return &out, nil
}
