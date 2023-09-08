package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode string

const (
	Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModeautomatic Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode = "Automatic"
	Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModedisabled  Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode = "Disabled"
	Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModemanual    Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode = "Manual"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModeautomatic),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModedisabled),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModemanual),
	}
}

func parseWindows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode(input string) (*Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode{
		"automatic": Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModeautomatic,
		"disabled":  Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModedisabled,
		"manual":    Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupModemanual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode(input)
	return &out, nil
}
