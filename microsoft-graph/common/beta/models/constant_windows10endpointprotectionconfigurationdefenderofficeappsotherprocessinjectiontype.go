package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeauditMode   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeblock       Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "Block"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypedisable     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "Disable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeuserDefined Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypewarn        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeblock),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypedisable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeauditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeblock,
		"disable":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypedisable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypeuserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType(input)
	return &out, nil
}
