package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviordisconnectRemoteDesktopSession Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "DisconnectRemoteDesktopSession"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviorforceLogoff                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "ForceLogoff"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviorlockWorkstation                Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "LockWorkstation"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviornoAction                       Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "NoAction"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviordisconnectRemoteDesktopSession),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviorforceLogoff),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviorlockWorkstation),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviornoAction),
	}
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior{
		"disconnectremotedesktopsession": Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviordisconnectRemoteDesktopSession,
		"forcelogoff":                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviorforceLogoff,
		"lockworkstation":                Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviorlockWorkstation,
		"noaction":                       Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehaviornoAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior(input)
	return &out, nil
}
