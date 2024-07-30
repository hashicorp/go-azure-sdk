package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType string

const (
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_AuditMode   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Block       Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "block"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Disable     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "disable"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_UserDefined Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Warn        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType(input string) (*Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScriptObfuscatedMacroCodeType(input)
	return &out, nil
}
