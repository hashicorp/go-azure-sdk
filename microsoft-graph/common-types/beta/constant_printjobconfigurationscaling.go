package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationScaling string

const (
	PrintJobConfigurationScaling_Auto        PrintJobConfigurationScaling = "auto"
	PrintJobConfigurationScaling_Fill        PrintJobConfigurationScaling = "fill"
	PrintJobConfigurationScaling_Fit         PrintJobConfigurationScaling = "fit"
	PrintJobConfigurationScaling_None        PrintJobConfigurationScaling = "none"
	PrintJobConfigurationScaling_ShrinkToFit PrintJobConfigurationScaling = "shrinkToFit"
)

func PossibleValuesForPrintJobConfigurationScaling() []string {
	return []string{
		string(PrintJobConfigurationScaling_Auto),
		string(PrintJobConfigurationScaling_Fill),
		string(PrintJobConfigurationScaling_Fit),
		string(PrintJobConfigurationScaling_None),
		string(PrintJobConfigurationScaling_ShrinkToFit),
	}
}

func (s *PrintJobConfigurationScaling) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationScaling(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationScaling(input string) (*PrintJobConfigurationScaling, error) {
	vals := map[string]PrintJobConfigurationScaling{
		"auto":        PrintJobConfigurationScaling_Auto,
		"fill":        PrintJobConfigurationScaling_Fill,
		"fit":         PrintJobConfigurationScaling_Fit,
		"none":        PrintJobConfigurationScaling_None,
		"shrinktofit": PrintJobConfigurationScaling_ShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationScaling(input)
	return &out, nil
}
