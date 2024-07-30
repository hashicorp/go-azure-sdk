package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsPrintColorConfiguration string

const (
	PrinterDefaultsPrintColorConfiguration_Auto          PrinterDefaultsPrintColorConfiguration = "auto"
	PrinterDefaultsPrintColorConfiguration_BlackAndWhite PrinterDefaultsPrintColorConfiguration = "blackAndWhite"
	PrinterDefaultsPrintColorConfiguration_Color         PrinterDefaultsPrintColorConfiguration = "color"
	PrinterDefaultsPrintColorConfiguration_Grayscale     PrinterDefaultsPrintColorConfiguration = "grayscale"
)

func PossibleValuesForPrinterDefaultsPrintColorConfiguration() []string {
	return []string{
		string(PrinterDefaultsPrintColorConfiguration_Auto),
		string(PrinterDefaultsPrintColorConfiguration_BlackAndWhite),
		string(PrinterDefaultsPrintColorConfiguration_Color),
		string(PrinterDefaultsPrintColorConfiguration_Grayscale),
	}
}

func (s *PrinterDefaultsPrintColorConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsPrintColorConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsPrintColorConfiguration(input string) (*PrinterDefaultsPrintColorConfiguration, error) {
	vals := map[string]PrinterDefaultsPrintColorConfiguration{
		"auto":          PrinterDefaultsPrintColorConfiguration_Auto,
		"blackandwhite": PrinterDefaultsPrintColorConfiguration_BlackAndWhite,
		"color":         PrinterDefaultsPrintColorConfiguration_Color,
		"grayscale":     PrinterDefaultsPrintColorConfiguration_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPrintColorConfiguration(input)
	return &out, nil
}
