package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay string

const (
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayInAppAndInNotifications Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "DisplayInAppAndInNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayOnlyInApp               Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "DisplayOnlyInApp"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayOnlyInNotifications     Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "DisplayOnlyInNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaynotConfigured                  Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayInAppAndInNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayOnlyInApp),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayOnlyInNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaynotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay(input string) (*Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay{
		"displayinappandinnotifications": Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayInAppAndInNotifications,
		"displayonlyinapp":               Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayOnlyInApp,
		"displayonlyinnotifications":     Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaydisplayOnlyInNotifications,
		"notconfigured":                  Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplaynotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderSecurityCenterITContactDisplay(input)
	return &out, nil
}
