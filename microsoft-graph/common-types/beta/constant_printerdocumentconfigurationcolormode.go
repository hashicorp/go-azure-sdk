package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationColorMode string

const (
	PrinterDocumentConfigurationColorMode_Auto          PrinterDocumentConfigurationColorMode = "auto"
	PrinterDocumentConfigurationColorMode_BlackAndWhite PrinterDocumentConfigurationColorMode = "blackAndWhite"
	PrinterDocumentConfigurationColorMode_Color         PrinterDocumentConfigurationColorMode = "color"
	PrinterDocumentConfigurationColorMode_Grayscale     PrinterDocumentConfigurationColorMode = "grayscale"
)

func PossibleValuesForPrinterDocumentConfigurationColorMode() []string {
	return []string{
		string(PrinterDocumentConfigurationColorMode_Auto),
		string(PrinterDocumentConfigurationColorMode_BlackAndWhite),
		string(PrinterDocumentConfigurationColorMode_Color),
		string(PrinterDocumentConfigurationColorMode_Grayscale),
	}
}

func (s *PrinterDocumentConfigurationColorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationColorMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationColorMode(input string) (*PrinterDocumentConfigurationColorMode, error) {
	vals := map[string]PrinterDocumentConfigurationColorMode{
		"auto":          PrinterDocumentConfigurationColorMode_Auto,
		"blackandwhite": PrinterDocumentConfigurationColorMode_BlackAndWhite,
		"color":         PrinterDocumentConfigurationColorMode_Color,
		"grayscale":     PrinterDocumentConfigurationColorMode_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationColorMode(input)
	return &out, nil
}
