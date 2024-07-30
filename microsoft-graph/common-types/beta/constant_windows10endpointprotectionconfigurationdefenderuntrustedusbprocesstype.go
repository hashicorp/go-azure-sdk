package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType string

const (
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_AuditMode   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Block       Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "block"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Disable     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "disable"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_UserDefined Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Warn        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType(input string) (*Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType(input)
	return &out, nil
}
