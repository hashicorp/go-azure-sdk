package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsColorMode string

const (
	PrinterDefaultsColorMode_Auto          PrinterDefaultsColorMode = "auto"
	PrinterDefaultsColorMode_BlackAndWhite PrinterDefaultsColorMode = "blackAndWhite"
	PrinterDefaultsColorMode_Color         PrinterDefaultsColorMode = "color"
	PrinterDefaultsColorMode_Grayscale     PrinterDefaultsColorMode = "grayscale"
)

func PossibleValuesForPrinterDefaultsColorMode() []string {
	return []string{
		string(PrinterDefaultsColorMode_Auto),
		string(PrinterDefaultsColorMode_BlackAndWhite),
		string(PrinterDefaultsColorMode_Color),
		string(PrinterDefaultsColorMode_Grayscale),
	}
}

func (s *PrinterDefaultsColorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsColorMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsColorMode(input string) (*PrinterDefaultsColorMode, error) {
	vals := map[string]PrinterDefaultsColorMode{
		"auto":          PrinterDefaultsColorMode_Auto,
		"blackandwhite": PrinterDefaultsColorMode_BlackAndWhite,
		"color":         PrinterDefaultsColorMode_Color,
		"grayscale":     PrinterDefaultsColorMode_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsColorMode(input)
	return &out, nil
}
