package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderProcessCreation string

const (
	Windows10EndpointProtectionConfigurationDefenderProcessCreation_AuditMode     Windows10EndpointProtectionConfigurationDefenderProcessCreation = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderProcessCreation_Enable        Windows10EndpointProtectionConfigurationDefenderProcessCreation = "enable"
	Windows10EndpointProtectionConfigurationDefenderProcessCreation_NotConfigured Windows10EndpointProtectionConfigurationDefenderProcessCreation = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderProcessCreation_UserDefined   Windows10EndpointProtectionConfigurationDefenderProcessCreation = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderProcessCreation_Warn          Windows10EndpointProtectionConfigurationDefenderProcessCreation = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderProcessCreation() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreation_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreation_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreation_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreation_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreation_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderProcessCreation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderProcessCreation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderProcessCreation(input string) (*Windows10EndpointProtectionConfigurationDefenderProcessCreation, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderProcessCreation{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderProcessCreation_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderProcessCreation_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderProcessCreation_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderProcessCreation_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderProcessCreation_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderProcessCreation(input)
	return &out, nil
}
