package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType string

const (
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeauditMode     Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeenable        Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "Enable"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypenotConfigured Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeuserDefined   Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypewarn          Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeenable),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypenotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType(input string) (*Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypenotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypeuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType(input)
	return &out, nil
}
