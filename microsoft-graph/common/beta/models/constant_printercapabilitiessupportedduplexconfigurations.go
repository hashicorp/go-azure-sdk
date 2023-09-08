package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedDuplexConfigurations string

const (
	PrinterCapabilitiesSupportedDuplexConfigurationsoneSided          PrinterCapabilitiesSupportedDuplexConfigurations = "OneSided"
	PrinterCapabilitiesSupportedDuplexConfigurationstwoSidedLongEdge  PrinterCapabilitiesSupportedDuplexConfigurations = "TwoSidedLongEdge"
	PrinterCapabilitiesSupportedDuplexConfigurationstwoSidedShortEdge PrinterCapabilitiesSupportedDuplexConfigurations = "TwoSidedShortEdge"
)

func PossibleValuesForPrinterCapabilitiesSupportedDuplexConfigurations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedDuplexConfigurationsoneSided),
		string(PrinterCapabilitiesSupportedDuplexConfigurationstwoSidedLongEdge),
		string(PrinterCapabilitiesSupportedDuplexConfigurationstwoSidedShortEdge),
	}
}

func parsePrinterCapabilitiesSupportedDuplexConfigurations(input string) (*PrinterCapabilitiesSupportedDuplexConfigurations, error) {
	vals := map[string]PrinterCapabilitiesSupportedDuplexConfigurations{
		"onesided":          PrinterCapabilitiesSupportedDuplexConfigurationsoneSided,
		"twosidedlongedge":  PrinterCapabilitiesSupportedDuplexConfigurationstwoSidedLongEdge,
		"twosidedshortedge": PrinterCapabilitiesSupportedDuplexConfigurationstwoSidedShortEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedDuplexConfigurations(input)
	return &out, nil
}
