package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_AuditMode   Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Block       Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType = "block"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Disable     Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType = "disable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_UserDefined Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Warn        Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunchType(input)
	return &out, nil
}
