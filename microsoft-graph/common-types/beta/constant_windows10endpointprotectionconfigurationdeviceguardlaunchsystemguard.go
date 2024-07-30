package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard string

const (
	Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_Disabled      Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard = "disabled"
	Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_Enabled       Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard = "enabled"
	Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_NotConfigured Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_Disabled),
		string(Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_Enabled),
		string(Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard(input string) (*Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard{
		"disabled":      Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_Disabled,
		"enabled":       Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_Enabled,
		"notconfigured": Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard(input)
	return &out, nil
}
