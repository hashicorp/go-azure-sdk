package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationDuplexMode string

const (
	PrintJobConfigurationDuplexMode_FlipOnLongEdge  PrintJobConfigurationDuplexMode = "flipOnLongEdge"
	PrintJobConfigurationDuplexMode_FlipOnShortEdge PrintJobConfigurationDuplexMode = "flipOnShortEdge"
	PrintJobConfigurationDuplexMode_OneSided        PrintJobConfigurationDuplexMode = "oneSided"
)

func PossibleValuesForPrintJobConfigurationDuplexMode() []string {
	return []string{
		string(PrintJobConfigurationDuplexMode_FlipOnLongEdge),
		string(PrintJobConfigurationDuplexMode_FlipOnShortEdge),
		string(PrintJobConfigurationDuplexMode_OneSided),
	}
}

func (s *PrintJobConfigurationDuplexMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationDuplexMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationDuplexMode(input string) (*PrintJobConfigurationDuplexMode, error) {
	vals := map[string]PrintJobConfigurationDuplexMode{
		"fliponlongedge":  PrintJobConfigurationDuplexMode_FlipOnLongEdge,
		"fliponshortedge": PrintJobConfigurationDuplexMode_FlipOnShortEdge,
		"onesided":        PrintJobConfigurationDuplexMode_OneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationDuplexMode(input)
	return &out, nil
}
