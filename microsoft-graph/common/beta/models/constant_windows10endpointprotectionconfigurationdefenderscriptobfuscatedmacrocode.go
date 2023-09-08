package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode string

const (
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeauditMode     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeenable        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "Enable"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodenotConfigured Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeuserDefined   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodewarn          Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeenable),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodenotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode(input string) (*Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodenotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode(input)
	return &out, nil
}
