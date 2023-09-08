package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType string

const (
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeauditMode     Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeenable        Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "Enable"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypenotConfigured Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeuserDefined   Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypewarn          Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeenable),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypenotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType(input string) (*Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypenotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypeuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType(input)
	return &out, nil
}
