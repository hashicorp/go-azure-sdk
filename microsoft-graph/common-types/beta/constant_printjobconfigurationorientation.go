package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationOrientation string

const (
	PrintJobConfigurationOrientation_Landscape        PrintJobConfigurationOrientation = "landscape"
	PrintJobConfigurationOrientation_Portrait         PrintJobConfigurationOrientation = "portrait"
	PrintJobConfigurationOrientation_ReverseLandscape PrintJobConfigurationOrientation = "reverseLandscape"
	PrintJobConfigurationOrientation_ReversePortrait  PrintJobConfigurationOrientation = "reversePortrait"
)

func PossibleValuesForPrintJobConfigurationOrientation() []string {
	return []string{
		string(PrintJobConfigurationOrientation_Landscape),
		string(PrintJobConfigurationOrientation_Portrait),
		string(PrintJobConfigurationOrientation_ReverseLandscape),
		string(PrintJobConfigurationOrientation_ReversePortrait),
	}
}

func (s *PrintJobConfigurationOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationOrientation(input string) (*PrintJobConfigurationOrientation, error) {
	vals := map[string]PrintJobConfigurationOrientation{
		"landscape":        PrintJobConfigurationOrientation_Landscape,
		"portrait":         PrintJobConfigurationOrientation_Portrait,
		"reverselandscape": PrintJobConfigurationOrientation_ReverseLandscape,
		"reverseportrait":  PrintJobConfigurationOrientation_ReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationOrientation(input)
	return &out, nil
}
