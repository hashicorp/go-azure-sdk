package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType string

const (
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeauditMode     Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeenable        Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "Enable"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypenotConfigured Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeuserDefined   Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypewarn          Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderNetworkProtectionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeenable),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypenotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderNetworkProtectionType(input string) (*Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypenotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypeuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderNetworkProtectionTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType(input)
	return &out, nil
}
