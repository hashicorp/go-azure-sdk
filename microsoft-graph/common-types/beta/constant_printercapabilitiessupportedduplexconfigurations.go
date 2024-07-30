package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedDuplexConfigurations string

const (
	PrinterCapabilitiesSupportedDuplexConfigurations_OneSided          PrinterCapabilitiesSupportedDuplexConfigurations = "oneSided"
	PrinterCapabilitiesSupportedDuplexConfigurations_TwoSidedLongEdge  PrinterCapabilitiesSupportedDuplexConfigurations = "twoSidedLongEdge"
	PrinterCapabilitiesSupportedDuplexConfigurations_TwoSidedShortEdge PrinterCapabilitiesSupportedDuplexConfigurations = "twoSidedShortEdge"
)

func PossibleValuesForPrinterCapabilitiesSupportedDuplexConfigurations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedDuplexConfigurations_OneSided),
		string(PrinterCapabilitiesSupportedDuplexConfigurations_TwoSidedLongEdge),
		string(PrinterCapabilitiesSupportedDuplexConfigurations_TwoSidedShortEdge),
	}
}

func (s *PrinterCapabilitiesSupportedDuplexConfigurations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedDuplexConfigurations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedDuplexConfigurations(input string) (*PrinterCapabilitiesSupportedDuplexConfigurations, error) {
	vals := map[string]PrinterCapabilitiesSupportedDuplexConfigurations{
		"onesided":          PrinterCapabilitiesSupportedDuplexConfigurations_OneSided,
		"twosidedlongedge":  PrinterCapabilitiesSupportedDuplexConfigurations_TwoSidedLongEdge,
		"twosidedshortedge": PrinterCapabilitiesSupportedDuplexConfigurations_TwoSidedShortEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedDuplexConfigurations(input)
	return &out, nil
}
