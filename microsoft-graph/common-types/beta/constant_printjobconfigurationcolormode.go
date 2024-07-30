package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationColorMode string

const (
	PrintJobConfigurationColorMode_Auto          PrintJobConfigurationColorMode = "auto"
	PrintJobConfigurationColorMode_BlackAndWhite PrintJobConfigurationColorMode = "blackAndWhite"
	PrintJobConfigurationColorMode_Color         PrintJobConfigurationColorMode = "color"
	PrintJobConfigurationColorMode_Grayscale     PrintJobConfigurationColorMode = "grayscale"
)

func PossibleValuesForPrintJobConfigurationColorMode() []string {
	return []string{
		string(PrintJobConfigurationColorMode_Auto),
		string(PrintJobConfigurationColorMode_BlackAndWhite),
		string(PrintJobConfigurationColorMode_Color),
		string(PrintJobConfigurationColorMode_Grayscale),
	}
}

func (s *PrintJobConfigurationColorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationColorMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationColorMode(input string) (*PrintJobConfigurationColorMode, error) {
	vals := map[string]PrintJobConfigurationColorMode{
		"auto":          PrintJobConfigurationColorMode_Auto,
		"blackandwhite": PrintJobConfigurationColorMode_BlackAndWhite,
		"color":         PrintJobConfigurationColorMode_Color,
		"grayscale":     PrintJobConfigurationColorMode_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationColorMode(input)
	return &out, nil
}
