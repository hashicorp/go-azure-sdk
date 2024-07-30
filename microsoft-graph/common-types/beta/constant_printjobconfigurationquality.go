package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationQuality string

const (
	PrintJobConfigurationQuality_High   PrintJobConfigurationQuality = "high"
	PrintJobConfigurationQuality_Low    PrintJobConfigurationQuality = "low"
	PrintJobConfigurationQuality_Medium PrintJobConfigurationQuality = "medium"
)

func PossibleValuesForPrintJobConfigurationQuality() []string {
	return []string{
		string(PrintJobConfigurationQuality_High),
		string(PrintJobConfigurationQuality_Low),
		string(PrintJobConfigurationQuality_Medium),
	}
}

func (s *PrintJobConfigurationQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationQuality(input string) (*PrintJobConfigurationQuality, error) {
	vals := map[string]PrintJobConfigurationQuality{
		"high":   PrintJobConfigurationQuality_High,
		"low":    PrintJobConfigurationQuality_Low,
		"medium": PrintJobConfigurationQuality_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationQuality(input)
	return &out, nil
}
