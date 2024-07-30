package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_AuditMode     Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_Enable        Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch = "enable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_NotConfigured Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_UserDefined   Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_Warn          Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsExecutableContentCreationOrLaunch(input)
	return &out, nil
}
