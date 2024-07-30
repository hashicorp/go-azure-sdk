package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType string

const (
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_AuditMode   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Block       Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "block"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Disable     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "disable"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_UserDefined Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Warn        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType(input string) (*Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderOfficeAppsOtherProcessInjectionType(input)
	return &out, nil
}
