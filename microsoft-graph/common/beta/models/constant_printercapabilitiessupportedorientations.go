package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedOrientations string

const (
	PrinterCapabilitiesSupportedOrientationslandscape        PrinterCapabilitiesSupportedOrientations = "Landscape"
	PrinterCapabilitiesSupportedOrientationsportrait         PrinterCapabilitiesSupportedOrientations = "Portrait"
	PrinterCapabilitiesSupportedOrientationsreverseLandscape PrinterCapabilitiesSupportedOrientations = "ReverseLandscape"
	PrinterCapabilitiesSupportedOrientationsreversePortrait  PrinterCapabilitiesSupportedOrientations = "ReversePortrait"
)

func PossibleValuesForPrinterCapabilitiesSupportedOrientations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedOrientationslandscape),
		string(PrinterCapabilitiesSupportedOrientationsportrait),
		string(PrinterCapabilitiesSupportedOrientationsreverseLandscape),
		string(PrinterCapabilitiesSupportedOrientationsreversePortrait),
	}
}

func parsePrinterCapabilitiesSupportedOrientations(input string) (*PrinterCapabilitiesSupportedOrientations, error) {
	vals := map[string]PrinterCapabilitiesSupportedOrientations{
		"landscape":        PrinterCapabilitiesSupportedOrientationslandscape,
		"portrait":         PrinterCapabilitiesSupportedOrientationsportrait,
		"reverselandscape": PrinterCapabilitiesSupportedOrientationsreverseLandscape,
		"reverseportrait":  PrinterCapabilitiesSupportedOrientationsreversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedOrientations(input)
	return &out, nil
}
