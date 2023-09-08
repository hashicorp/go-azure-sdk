package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesOrientations string

const (
	PrinterCapabilitiesOrientationslandscape        PrinterCapabilitiesOrientations = "Landscape"
	PrinterCapabilitiesOrientationsportrait         PrinterCapabilitiesOrientations = "Portrait"
	PrinterCapabilitiesOrientationsreverseLandscape PrinterCapabilitiesOrientations = "ReverseLandscape"
	PrinterCapabilitiesOrientationsreversePortrait  PrinterCapabilitiesOrientations = "ReversePortrait"
)

func PossibleValuesForPrinterCapabilitiesOrientations() []string {
	return []string{
		string(PrinterCapabilitiesOrientationslandscape),
		string(PrinterCapabilitiesOrientationsportrait),
		string(PrinterCapabilitiesOrientationsreverseLandscape),
		string(PrinterCapabilitiesOrientationsreversePortrait),
	}
}

func parsePrinterCapabilitiesOrientations(input string) (*PrinterCapabilitiesOrientations, error) {
	vals := map[string]PrinterCapabilitiesOrientations{
		"landscape":        PrinterCapabilitiesOrientationslandscape,
		"portrait":         PrinterCapabilitiesOrientationsportrait,
		"reverselandscape": PrinterCapabilitiesOrientationsreverseLandscape,
		"reverseportrait":  PrinterCapabilitiesOrientationsreversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesOrientations(input)
	return &out, nil
}
