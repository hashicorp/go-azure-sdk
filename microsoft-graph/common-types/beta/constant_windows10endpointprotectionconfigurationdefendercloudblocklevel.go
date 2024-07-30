package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel string

const (
	Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_High          Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel = "high"
	Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_HighPlus      Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel = "highPlus"
	Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_NotConfigured Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_ZeroTolerance Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel = "zeroTolerance"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderCloudBlockLevel() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_High),
		string(Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_HighPlus),
		string(Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_ZeroTolerance),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderCloudBlockLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderCloudBlockLevel(input string) (*Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel{
		"high":          Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_High,
		"highplus":      Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_HighPlus,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_NotConfigured,
		"zerotolerance": Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel_ZeroTolerance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderCloudBlockLevel(input)
	return &out, nil
}
