package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsColorMode string

const (
	PrinterDefaultsColorModeauto          PrinterDefaultsColorMode = "Auto"
	PrinterDefaultsColorModeblackAndWhite PrinterDefaultsColorMode = "BlackAndWhite"
	PrinterDefaultsColorModecolor         PrinterDefaultsColorMode = "Color"
	PrinterDefaultsColorModegrayscale     PrinterDefaultsColorMode = "Grayscale"
)

func PossibleValuesForPrinterDefaultsColorMode() []string {
	return []string{
		string(PrinterDefaultsColorModeauto),
		string(PrinterDefaultsColorModeblackAndWhite),
		string(PrinterDefaultsColorModecolor),
		string(PrinterDefaultsColorModegrayscale),
	}
}

func parsePrinterDefaultsColorMode(input string) (*PrinterDefaultsColorMode, error) {
	vals := map[string]PrinterDefaultsColorMode{
		"auto":          PrinterDefaultsColorModeauto,
		"blackandwhite": PrinterDefaultsColorModeblackAndWhite,
		"color":         PrinterDefaultsColorModecolor,
		"grayscale":     PrinterDefaultsColorModegrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsColorMode(input)
	return &out, nil
}
