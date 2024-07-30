package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsScaling string

const (
	PrinterDefaultsScaling_Auto        PrinterDefaultsScaling = "auto"
	PrinterDefaultsScaling_Fill        PrinterDefaultsScaling = "fill"
	PrinterDefaultsScaling_Fit         PrinterDefaultsScaling = "fit"
	PrinterDefaultsScaling_None        PrinterDefaultsScaling = "none"
	PrinterDefaultsScaling_ShrinkToFit PrinterDefaultsScaling = "shrinkToFit"
)

func PossibleValuesForPrinterDefaultsScaling() []string {
	return []string{
		string(PrinterDefaultsScaling_Auto),
		string(PrinterDefaultsScaling_Fill),
		string(PrinterDefaultsScaling_Fit),
		string(PrinterDefaultsScaling_None),
		string(PrinterDefaultsScaling_ShrinkToFit),
	}
}

func (s *PrinterDefaultsScaling) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsScaling(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsScaling(input string) (*PrinterDefaultsScaling, error) {
	vals := map[string]PrinterDefaultsScaling{
		"auto":        PrinterDefaultsScaling_Auto,
		"fill":        PrinterDefaultsScaling_Fill,
		"fit":         PrinterDefaultsScaling_Fit,
		"none":        PrinterDefaultsScaling_None,
		"shrinktofit": PrinterDefaultsScaling_ShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsScaling(input)
	return &out, nil
}
