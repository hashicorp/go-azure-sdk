package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_AuditMode     Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_Enable        Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess = "enable"
	Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_NotConfigured Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_UserDefined   Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_Warn          Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeCommunicationAppsLaunchChildProcess(input)
	return &out, nil
}
