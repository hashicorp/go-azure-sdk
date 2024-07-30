package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess string

const (
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_AuditMode     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_Enable        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "enable"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_NotConfigured Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_UserDefined   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_Warn          Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess(input string) (*Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess(input)
	return &out, nil
}
