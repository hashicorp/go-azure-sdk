package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesScalings string

const (
	PrinterCapabilitiesScalings_Auto        PrinterCapabilitiesScalings = "auto"
	PrinterCapabilitiesScalings_Fill        PrinterCapabilitiesScalings = "fill"
	PrinterCapabilitiesScalings_Fit         PrinterCapabilitiesScalings = "fit"
	PrinterCapabilitiesScalings_None        PrinterCapabilitiesScalings = "none"
	PrinterCapabilitiesScalings_ShrinkToFit PrinterCapabilitiesScalings = "shrinkToFit"
)

func PossibleValuesForPrinterCapabilitiesScalings() []string {
	return []string{
		string(PrinterCapabilitiesScalings_Auto),
		string(PrinterCapabilitiesScalings_Fill),
		string(PrinterCapabilitiesScalings_Fit),
		string(PrinterCapabilitiesScalings_None),
		string(PrinterCapabilitiesScalings_ShrinkToFit),
	}
}

func (s *PrinterCapabilitiesScalings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesScalings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesScalings(input string) (*PrinterCapabilitiesScalings, error) {
	vals := map[string]PrinterCapabilitiesScalings{
		"auto":        PrinterCapabilitiesScalings_Auto,
		"fill":        PrinterCapabilitiesScalings_Fill,
		"fit":         PrinterCapabilitiesScalings_Fit,
		"none":        PrinterCapabilitiesScalings_None,
		"shrinktofit": PrinterCapabilitiesScalings_ShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesScalings(input)
	return &out, nil
}
