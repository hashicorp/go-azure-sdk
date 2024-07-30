package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchExportOperationExportLocation string

const (
	SecurityEdiscoverySearchExportOperationExportLocation_NonresponsiveLocations SecurityEdiscoverySearchExportOperationExportLocation = "nonresponsiveLocations"
	SecurityEdiscoverySearchExportOperationExportLocation_ResponsiveLocations    SecurityEdiscoverySearchExportOperationExportLocation = "responsiveLocations"
)

func PossibleValuesForSecurityEdiscoverySearchExportOperationExportLocation() []string {
	return []string{
		string(SecurityEdiscoverySearchExportOperationExportLocation_NonresponsiveLocations),
		string(SecurityEdiscoverySearchExportOperationExportLocation_ResponsiveLocations),
	}
}

func (s *SecurityEdiscoverySearchExportOperationExportLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoverySearchExportOperationExportLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoverySearchExportOperationExportLocation(input string) (*SecurityEdiscoverySearchExportOperationExportLocation, error) {
	vals := map[string]SecurityEdiscoverySearchExportOperationExportLocation{
		"nonresponsivelocations": SecurityEdiscoverySearchExportOperationExportLocation_NonresponsiveLocations,
		"responsivelocations":    SecurityEdiscoverySearchExportOperationExportLocation_ResponsiveLocations,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchExportOperationExportLocation(input)
	return &out, nil
}
