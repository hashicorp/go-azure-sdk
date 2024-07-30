package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_AuditMode   Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Block       Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType = "block"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Disable     Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType = "disable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_UserDefined Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Warn        Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsLaunchChildProcessType(input)
	return &out, nil
}
