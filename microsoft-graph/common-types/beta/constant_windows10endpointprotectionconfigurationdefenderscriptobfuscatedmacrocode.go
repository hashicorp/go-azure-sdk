package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode string

const (
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_AuditMode     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_Enable        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "enable"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_NotConfigured Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_UserDefined   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_Warn          Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode(input string) (*Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCode(input)
	return &out, nil
}
