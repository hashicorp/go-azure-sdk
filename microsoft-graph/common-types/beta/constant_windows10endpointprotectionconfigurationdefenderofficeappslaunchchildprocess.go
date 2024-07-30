package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_AuditMode     Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_Enable        Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess = "enable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_NotConfigured Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_UserDefined   Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_Warn          Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcess(input)
	return &out, nil
}
