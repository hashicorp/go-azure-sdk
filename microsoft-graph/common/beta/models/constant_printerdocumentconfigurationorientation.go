package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationOrientation string

const (
	PrinterDocumentConfigurationOrientationlandscape        PrinterDocumentConfigurationOrientation = "Landscape"
	PrinterDocumentConfigurationOrientationportrait         PrinterDocumentConfigurationOrientation = "Portrait"
	PrinterDocumentConfigurationOrientationreverseLandscape PrinterDocumentConfigurationOrientation = "ReverseLandscape"
	PrinterDocumentConfigurationOrientationreversePortrait  PrinterDocumentConfigurationOrientation = "ReversePortrait"
)

func PossibleValuesForPrinterDocumentConfigurationOrientation() []string {
	return []string{
		string(PrinterDocumentConfigurationOrientationlandscape),
		string(PrinterDocumentConfigurationOrientationportrait),
		string(PrinterDocumentConfigurationOrientationreverseLandscape),
		string(PrinterDocumentConfigurationOrientationreversePortrait),
	}
}

func parsePrinterDocumentConfigurationOrientation(input string) (*PrinterDocumentConfigurationOrientation, error) {
	vals := map[string]PrinterDocumentConfigurationOrientation{
		"landscape":        PrinterDocumentConfigurationOrientationlandscape,
		"portrait":         PrinterDocumentConfigurationOrientationportrait,
		"reverselandscape": PrinterDocumentConfigurationOrientationreverseLandscape,
		"reverseportrait":  PrinterDocumentConfigurationOrientationreversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationOrientation(input)
	return &out, nil
}
