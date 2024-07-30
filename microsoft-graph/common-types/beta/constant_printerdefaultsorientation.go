package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsOrientation string

const (
	PrinterDefaultsOrientation_Landscape        PrinterDefaultsOrientation = "landscape"
	PrinterDefaultsOrientation_Portrait         PrinterDefaultsOrientation = "portrait"
	PrinterDefaultsOrientation_ReverseLandscape PrinterDefaultsOrientation = "reverseLandscape"
	PrinterDefaultsOrientation_ReversePortrait  PrinterDefaultsOrientation = "reversePortrait"
)

func PossibleValuesForPrinterDefaultsOrientation() []string {
	return []string{
		string(PrinterDefaultsOrientation_Landscape),
		string(PrinterDefaultsOrientation_Portrait),
		string(PrinterDefaultsOrientation_ReverseLandscape),
		string(PrinterDefaultsOrientation_ReversePortrait),
	}
}

func (s *PrinterDefaultsOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsOrientation(input string) (*PrinterDefaultsOrientation, error) {
	vals := map[string]PrinterDefaultsOrientation{
		"landscape":        PrinterDefaultsOrientation_Landscape,
		"portrait":         PrinterDefaultsOrientation_Portrait,
		"reverselandscape": PrinterDefaultsOrientation_ReverseLandscape,
		"reverseportrait":  PrinterDefaultsOrientation_ReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsOrientation(input)
	return &out, nil
}
