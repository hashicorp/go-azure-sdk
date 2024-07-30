package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesOrientations string

const (
	PrinterCapabilitiesOrientations_Landscape        PrinterCapabilitiesOrientations = "landscape"
	PrinterCapabilitiesOrientations_Portrait         PrinterCapabilitiesOrientations = "portrait"
	PrinterCapabilitiesOrientations_ReverseLandscape PrinterCapabilitiesOrientations = "reverseLandscape"
	PrinterCapabilitiesOrientations_ReversePortrait  PrinterCapabilitiesOrientations = "reversePortrait"
)

func PossibleValuesForPrinterCapabilitiesOrientations() []string {
	return []string{
		string(PrinterCapabilitiesOrientations_Landscape),
		string(PrinterCapabilitiesOrientations_Portrait),
		string(PrinterCapabilitiesOrientations_ReverseLandscape),
		string(PrinterCapabilitiesOrientations_ReversePortrait),
	}
}

func (s *PrinterCapabilitiesOrientations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesOrientations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesOrientations(input string) (*PrinterCapabilitiesOrientations, error) {
	vals := map[string]PrinterCapabilitiesOrientations{
		"landscape":        PrinterCapabilitiesOrientations_Landscape,
		"portrait":         PrinterCapabilitiesOrientations_Portrait,
		"reverselandscape": PrinterCapabilitiesOrientations_ReverseLandscape,
		"reverseportrait":  PrinterCapabilitiesOrientations_ReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesOrientations(input)
	return &out, nil
}
