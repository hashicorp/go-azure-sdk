package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay string

const (
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayInAppAndInNotifications Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "displayInAppAndInNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayOnlyInApp               Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "displayOnlyInApp"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayOnlyInNotifications     Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "displayOnlyInNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_NotConfigured                  Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayInAppAndInNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayOnlyInApp),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayOnlyInNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay(input string) (*Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay{
		"displayinappandinnotifications": Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayInAppAndInNotifications,
		"displayonlyinapp":               Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayOnlyInApp,
		"displayonlyinnotifications":     Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_DisplayOnlyInNotifications,
		"notconfigured":                  Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay(input)
	return &out, nil
}
