package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_AuditMode     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_Enable        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "enable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_NotConfigured Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_UserDefined   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_Warn          Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjection(input)
	return &out, nil
}
