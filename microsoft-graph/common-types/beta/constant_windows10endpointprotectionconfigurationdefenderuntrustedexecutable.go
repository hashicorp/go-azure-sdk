package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable string

const (
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_AuditMode     Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_Enable        Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable = "enable"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_NotConfigured Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_UserDefined   Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_Warn          Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderUntrustedExecutable() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderUntrustedExecutable(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderUntrustedExecutable(input string) (*Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderUntrustedExecutable(input)
	return &out, nil
}
