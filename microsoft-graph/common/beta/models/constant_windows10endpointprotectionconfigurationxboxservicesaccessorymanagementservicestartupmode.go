package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode string

const (
	Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModeautomatic Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode = "Automatic"
	Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModedisabled  Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode = "Disabled"
	Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModemanual    Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode = "Manual"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModeautomatic),
		string(Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModedisabled),
		string(Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModemanual),
	}
}

func parseWindows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode(input string) (*Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode{
		"automatic": Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModeautomatic,
		"disabled":  Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModedisabled,
		"manual":    Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupModemanual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode(input)
	return &out, nil
}
