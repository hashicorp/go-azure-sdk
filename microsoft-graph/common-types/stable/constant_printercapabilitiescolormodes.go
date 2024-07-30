package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesColorModes string

const (
	PrinterCapabilitiesColorModes_Auto          PrinterCapabilitiesColorModes = "auto"
	PrinterCapabilitiesColorModes_BlackAndWhite PrinterCapabilitiesColorModes = "blackAndWhite"
	PrinterCapabilitiesColorModes_Color         PrinterCapabilitiesColorModes = "color"
	PrinterCapabilitiesColorModes_Grayscale     PrinterCapabilitiesColorModes = "grayscale"
)

func PossibleValuesForPrinterCapabilitiesColorModes() []string {
	return []string{
		string(PrinterCapabilitiesColorModes_Auto),
		string(PrinterCapabilitiesColorModes_BlackAndWhite),
		string(PrinterCapabilitiesColorModes_Color),
		string(PrinterCapabilitiesColorModes_Grayscale),
	}
}

func (s *PrinterCapabilitiesColorModes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesColorModes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesColorModes(input string) (*PrinterCapabilitiesColorModes, error) {
	vals := map[string]PrinterCapabilitiesColorModes{
		"auto":          PrinterCapabilitiesColorModes_Auto,
		"blackandwhite": PrinterCapabilitiesColorModes_BlackAndWhite,
		"color":         PrinterCapabilitiesColorModes_Color,
		"grayscale":     PrinterCapabilitiesColorModes_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesColorModes(input)
	return &out, nil
}
