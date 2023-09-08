package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction string

const (
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionauditMode     Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionenable        Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "Enable"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionnotConfigured Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionuserDefined   Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionwarn          Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionenable),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionnotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionwarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction(input string) (*Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionnotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppActionwarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction(input)
	return &out, nil
}
