package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionauditMode     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionenable        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "Enable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionnotConfigured Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionuserDefined   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionwarn          Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionenable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionnotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionwarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionnotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionwarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection(input)
	return &out, nil
}
