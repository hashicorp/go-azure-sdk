package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType string

const (
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeauditMode   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeblock       Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "Block"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypedisable     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "Disable"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeuserDefined Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypewarn        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeblock),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypedisable),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType(input string) (*Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeauditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeblock,
		"disable":     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypedisable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypeuserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType(input)
	return &out, nil
}
