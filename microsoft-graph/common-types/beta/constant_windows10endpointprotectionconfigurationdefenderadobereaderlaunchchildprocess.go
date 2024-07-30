package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess string

const (
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_AuditMode     Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_Enable        Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "enable"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_NotConfigured Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_UserDefined   Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_Warn          Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess(input string) (*Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess(input)
	return &out, nil
}
