package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType string

const (
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_AuditMode   Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Block       Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType = "block"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Disable     Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType = "disable"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_UserDefined Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Warn        Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderUntrustedExecutableType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderUntrustedExecutableType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderUntrustedExecutableType(input string) (*Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderUntrustedExecutableType(input)
	return &out, nil
}
