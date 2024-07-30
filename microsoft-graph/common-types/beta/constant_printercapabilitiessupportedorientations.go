package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedOrientations string

const (
	PrinterCapabilitiesSupportedOrientations_Landscape        PrinterCapabilitiesSupportedOrientations = "landscape"
	PrinterCapabilitiesSupportedOrientations_Portrait         PrinterCapabilitiesSupportedOrientations = "portrait"
	PrinterCapabilitiesSupportedOrientations_ReverseLandscape PrinterCapabilitiesSupportedOrientations = "reverseLandscape"
	PrinterCapabilitiesSupportedOrientations_ReversePortrait  PrinterCapabilitiesSupportedOrientations = "reversePortrait"
)

func PossibleValuesForPrinterCapabilitiesSupportedOrientations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedOrientations_Landscape),
		string(PrinterCapabilitiesSupportedOrientations_Portrait),
		string(PrinterCapabilitiesSupportedOrientations_ReverseLandscape),
		string(PrinterCapabilitiesSupportedOrientations_ReversePortrait),
	}
}

func (s *PrinterCapabilitiesSupportedOrientations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedOrientations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedOrientations(input string) (*PrinterCapabilitiesSupportedOrientations, error) {
	vals := map[string]PrinterCapabilitiesSupportedOrientations{
		"landscape":        PrinterCapabilitiesSupportedOrientations_Landscape,
		"portrait":         PrinterCapabilitiesSupportedOrientations_Portrait,
		"reverselandscape": PrinterCapabilitiesSupportedOrientations_ReverseLandscape,
		"reverseportrait":  PrinterCapabilitiesSupportedOrientations_ReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedOrientations(input)
	return &out, nil
}
