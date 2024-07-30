package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationScaling string

const (
	PrinterDocumentConfigurationScaling_Auto        PrinterDocumentConfigurationScaling = "auto"
	PrinterDocumentConfigurationScaling_Fill        PrinterDocumentConfigurationScaling = "fill"
	PrinterDocumentConfigurationScaling_Fit         PrinterDocumentConfigurationScaling = "fit"
	PrinterDocumentConfigurationScaling_None        PrinterDocumentConfigurationScaling = "none"
	PrinterDocumentConfigurationScaling_ShrinkToFit PrinterDocumentConfigurationScaling = "shrinkToFit"
)

func PossibleValuesForPrinterDocumentConfigurationScaling() []string {
	return []string{
		string(PrinterDocumentConfigurationScaling_Auto),
		string(PrinterDocumentConfigurationScaling_Fill),
		string(PrinterDocumentConfigurationScaling_Fit),
		string(PrinterDocumentConfigurationScaling_None),
		string(PrinterDocumentConfigurationScaling_ShrinkToFit),
	}
}

func (s *PrinterDocumentConfigurationScaling) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationScaling(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationScaling(input string) (*PrinterDocumentConfigurationScaling, error) {
	vals := map[string]PrinterDocumentConfigurationScaling{
		"auto":        PrinterDocumentConfigurationScaling_Auto,
		"fill":        PrinterDocumentConfigurationScaling_Fill,
		"fit":         PrinterDocumentConfigurationScaling_Fit,
		"none":        PrinterDocumentConfigurationScaling_None,
		"shrinktofit": PrinterDocumentConfigurationScaling_ShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationScaling(input)
	return &out, nil
}
