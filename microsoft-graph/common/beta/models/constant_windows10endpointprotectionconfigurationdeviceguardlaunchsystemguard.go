package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard string

const (
	Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuarddisabled      Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard = "Disabled"
	Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuardenabled       Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard = "Enabled"
	Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuardnotConfigured Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuarddisabled),
		string(Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuardenabled),
		string(Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuardnotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard(input string) (*Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard{
		"disabled":      Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuarddisabled,
		"enabled":       Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuardenabled,
		"notconfigured": Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuardnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDeviceGuardLaunchSystemGuard(input)
	return &out, nil
}
