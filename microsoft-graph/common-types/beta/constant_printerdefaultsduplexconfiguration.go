package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsDuplexConfiguration string

const (
	PrinterDefaultsDuplexConfiguration_OneSided          PrinterDefaultsDuplexConfiguration = "oneSided"
	PrinterDefaultsDuplexConfiguration_TwoSidedLongEdge  PrinterDefaultsDuplexConfiguration = "twoSidedLongEdge"
	PrinterDefaultsDuplexConfiguration_TwoSidedShortEdge PrinterDefaultsDuplexConfiguration = "twoSidedShortEdge"
)

func PossibleValuesForPrinterDefaultsDuplexConfiguration() []string {
	return []string{
		string(PrinterDefaultsDuplexConfiguration_OneSided),
		string(PrinterDefaultsDuplexConfiguration_TwoSidedLongEdge),
		string(PrinterDefaultsDuplexConfiguration_TwoSidedShortEdge),
	}
}

func (s *PrinterDefaultsDuplexConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsDuplexConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsDuplexConfiguration(input string) (*PrinterDefaultsDuplexConfiguration, error) {
	vals := map[string]PrinterDefaultsDuplexConfiguration{
		"onesided":          PrinterDefaultsDuplexConfiguration_OneSided,
		"twosidedlongedge":  PrinterDefaultsDuplexConfiguration_TwoSidedLongEdge,
		"twosidedshortedge": PrinterDefaultsDuplexConfiguration_TwoSidedShortEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsDuplexConfiguration(input)
	return &out, nil
}
