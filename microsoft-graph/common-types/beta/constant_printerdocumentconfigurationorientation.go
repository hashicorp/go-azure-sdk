package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationOrientation string

const (
	PrinterDocumentConfigurationOrientation_Landscape        PrinterDocumentConfigurationOrientation = "landscape"
	PrinterDocumentConfigurationOrientation_Portrait         PrinterDocumentConfigurationOrientation = "portrait"
	PrinterDocumentConfigurationOrientation_ReverseLandscape PrinterDocumentConfigurationOrientation = "reverseLandscape"
	PrinterDocumentConfigurationOrientation_ReversePortrait  PrinterDocumentConfigurationOrientation = "reversePortrait"
)

func PossibleValuesForPrinterDocumentConfigurationOrientation() []string {
	return []string{
		string(PrinterDocumentConfigurationOrientation_Landscape),
		string(PrinterDocumentConfigurationOrientation_Portrait),
		string(PrinterDocumentConfigurationOrientation_ReverseLandscape),
		string(PrinterDocumentConfigurationOrientation_ReversePortrait),
	}
}

func (s *PrinterDocumentConfigurationOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationOrientation(input string) (*PrinterDocumentConfigurationOrientation, error) {
	vals := map[string]PrinterDocumentConfigurationOrientation{
		"landscape":        PrinterDocumentConfigurationOrientation_Landscape,
		"portrait":         PrinterDocumentConfigurationOrientation_Portrait,
		"reverselandscape": PrinterDocumentConfigurationOrientation_ReverseLandscape,
		"reverseportrait":  PrinterDocumentConfigurationOrientation_ReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationOrientation(input)
	return &out, nil
}
