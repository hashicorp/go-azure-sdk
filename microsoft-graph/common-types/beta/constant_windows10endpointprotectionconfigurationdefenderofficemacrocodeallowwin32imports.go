package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_AuditMode     Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_Enable        Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports = "enable"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_NotConfigured Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_UserDefined   Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_Warn          Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeMacroCodeAllowWin32Imports(input)
	return &out, nil
}
