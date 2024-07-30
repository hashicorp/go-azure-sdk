package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedColorConfigurations string

const (
	PrinterCapabilitiesSupportedColorConfigurations_Auto          PrinterCapabilitiesSupportedColorConfigurations = "auto"
	PrinterCapabilitiesSupportedColorConfigurations_BlackAndWhite PrinterCapabilitiesSupportedColorConfigurations = "blackAndWhite"
	PrinterCapabilitiesSupportedColorConfigurations_Color         PrinterCapabilitiesSupportedColorConfigurations = "color"
	PrinterCapabilitiesSupportedColorConfigurations_Grayscale     PrinterCapabilitiesSupportedColorConfigurations = "grayscale"
)

func PossibleValuesForPrinterCapabilitiesSupportedColorConfigurations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedColorConfigurations_Auto),
		string(PrinterCapabilitiesSupportedColorConfigurations_BlackAndWhite),
		string(PrinterCapabilitiesSupportedColorConfigurations_Color),
		string(PrinterCapabilitiesSupportedColorConfigurations_Grayscale),
	}
}

func (s *PrinterCapabilitiesSupportedColorConfigurations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedColorConfigurations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedColorConfigurations(input string) (*PrinterCapabilitiesSupportedColorConfigurations, error) {
	vals := map[string]PrinterCapabilitiesSupportedColorConfigurations{
		"auto":          PrinterCapabilitiesSupportedColorConfigurations_Auto,
		"blackandwhite": PrinterCapabilitiesSupportedColorConfigurations_BlackAndWhite,
		"color":         PrinterCapabilitiesSupportedColorConfigurations_Color,
		"grayscale":     PrinterCapabilitiesSupportedColorConfigurations_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedColorConfigurations(input)
	return &out, nil
}
