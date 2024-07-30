package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_AuditMode   Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Block       Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "block"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Disable     Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "disable"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_UserDefined Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Warn        Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32ImportsType(input)
	return &out, nil
}
