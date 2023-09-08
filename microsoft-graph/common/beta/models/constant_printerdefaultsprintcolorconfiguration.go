package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsPrintColorConfiguration string

const (
	PrinterDefaultsPrintColorConfigurationauto          PrinterDefaultsPrintColorConfiguration = "Auto"
	PrinterDefaultsPrintColorConfigurationblackAndWhite PrinterDefaultsPrintColorConfiguration = "BlackAndWhite"
	PrinterDefaultsPrintColorConfigurationcolor         PrinterDefaultsPrintColorConfiguration = "Color"
	PrinterDefaultsPrintColorConfigurationgrayscale     PrinterDefaultsPrintColorConfiguration = "Grayscale"
)

func PossibleValuesForPrinterDefaultsPrintColorConfiguration() []string {
	return []string{
		string(PrinterDefaultsPrintColorConfigurationauto),
		string(PrinterDefaultsPrintColorConfigurationblackAndWhite),
		string(PrinterDefaultsPrintColorConfigurationcolor),
		string(PrinterDefaultsPrintColorConfigurationgrayscale),
	}
}

func parsePrinterDefaultsPrintColorConfiguration(input string) (*PrinterDefaultsPrintColorConfiguration, error) {
	vals := map[string]PrinterDefaultsPrintColorConfiguration{
		"auto":          PrinterDefaultsPrintColorConfigurationauto,
		"blackandwhite": PrinterDefaultsPrintColorConfigurationblackAndWhite,
		"color":         PrinterDefaultsPrintColorConfigurationcolor,
		"grayscale":     PrinterDefaultsPrintColorConfigurationgrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPrintColorConfiguration(input)
	return &out, nil
}
