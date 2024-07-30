package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScanType string

const (
	Windows10EndpointProtectionConfigurationDefenderScanType_Disabled    Windows10EndpointProtectionConfigurationDefenderScanType = "disabled"
	Windows10EndpointProtectionConfigurationDefenderScanType_Full        Windows10EndpointProtectionConfigurationDefenderScanType = "full"
	Windows10EndpointProtectionConfigurationDefenderScanType_Quick       Windows10EndpointProtectionConfigurationDefenderScanType = "quick"
	Windows10EndpointProtectionConfigurationDefenderScanType_UserDefined Windows10EndpointProtectionConfigurationDefenderScanType = "userDefined"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScanType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScanType_Disabled),
		string(Windows10EndpointProtectionConfigurationDefenderScanType_Full),
		string(Windows10EndpointProtectionConfigurationDefenderScanType_Quick),
		string(Windows10EndpointProtectionConfigurationDefenderScanType_UserDefined),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderScanType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderScanType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderScanType(input string) (*Windows10EndpointProtectionConfigurationDefenderScanType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScanType{
		"disabled":    Windows10EndpointProtectionConfigurationDefenderScanType_Disabled,
		"full":        Windows10EndpointProtectionConfigurationDefenderScanType_Full,
		"quick":       Windows10EndpointProtectionConfigurationDefenderScanType_Quick,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderScanType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScanType(input)
	return &out, nil
}
