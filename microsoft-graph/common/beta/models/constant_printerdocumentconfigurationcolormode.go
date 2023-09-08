package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationColorMode string

const (
	PrinterDocumentConfigurationColorModeauto          PrinterDocumentConfigurationColorMode = "Auto"
	PrinterDocumentConfigurationColorModeblackAndWhite PrinterDocumentConfigurationColorMode = "BlackAndWhite"
	PrinterDocumentConfigurationColorModecolor         PrinterDocumentConfigurationColorMode = "Color"
	PrinterDocumentConfigurationColorModegrayscale     PrinterDocumentConfigurationColorMode = "Grayscale"
)

func PossibleValuesForPrinterDocumentConfigurationColorMode() []string {
	return []string{
		string(PrinterDocumentConfigurationColorModeauto),
		string(PrinterDocumentConfigurationColorModeblackAndWhite),
		string(PrinterDocumentConfigurationColorModecolor),
		string(PrinterDocumentConfigurationColorModegrayscale),
	}
}

func parsePrinterDocumentConfigurationColorMode(input string) (*PrinterDocumentConfigurationColorMode, error) {
	vals := map[string]PrinterDocumentConfigurationColorMode{
		"auto":          PrinterDocumentConfigurationColorModeauto,
		"blackandwhite": PrinterDocumentConfigurationColorModeblackAndWhite,
		"color":         PrinterDocumentConfigurationColorModecolor,
		"grayscale":     PrinterDocumentConfigurationColorModegrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationColorMode(input)
	return &out, nil
}
