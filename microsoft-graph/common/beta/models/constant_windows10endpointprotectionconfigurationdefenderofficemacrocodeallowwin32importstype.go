package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeauditMode   Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeblock       Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "Block"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypedisable     Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "Disable"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeuserDefined Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypewarn        Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeblock),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypedisable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeauditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeblock,
		"disable":     Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypedisable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypeuserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType(input)
	return &out, nil
}
