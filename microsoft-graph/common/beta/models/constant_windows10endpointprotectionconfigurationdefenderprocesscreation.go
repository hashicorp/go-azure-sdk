package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderProcessCreation string

const (
	Windows10EndpointProtectionConfigurationDefenderProcessCreationauditMode     Windows10EndpointProtectionConfigurationDefenderProcessCreation = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationenable        Windows10EndpointProtectionConfigurationDefenderProcessCreation = "Enable"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationnotConfigured Windows10EndpointProtectionConfigurationDefenderProcessCreation = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationuserDefined   Windows10EndpointProtectionConfigurationDefenderProcessCreation = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationwarn          Windows10EndpointProtectionConfigurationDefenderProcessCreation = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderProcessCreation() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationenable),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationnotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationwarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderProcessCreation(input string) (*Windows10EndpointProtectionConfigurationDefenderProcessCreation, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderProcessCreation{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderProcessCreationauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderProcessCreationenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderProcessCreationnotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderProcessCreationuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderProcessCreationwarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderProcessCreation(input)
	return &out, nil
}
