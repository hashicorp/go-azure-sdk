package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp string

const (
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppblockAllNotifications         Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp = "BlockAllNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppblockNoncriticalNotifications Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp = "BlockNoncriticalNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppnotConfigured                 Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppblockAllNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppblockNoncriticalNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppnotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp(input string) (*Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp{
		"blockallnotifications":         Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppblockAllNotifications,
		"blocknoncriticalnotifications": Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppblockNoncriticalNotifications,
		"notconfigured":                 Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromAppnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp(input)
	return &out, nil
}
