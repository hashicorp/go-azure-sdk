package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction string

const (
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_AuditMode     Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_Enable        Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "enable"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_NotConfigured Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_UserDefined   Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_Warn          Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction(input string) (*Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderPotentiallyUnwantedAppAction(input)
	return &out, nil
}
