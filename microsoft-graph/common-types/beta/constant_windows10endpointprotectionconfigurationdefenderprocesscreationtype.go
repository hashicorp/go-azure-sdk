package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderProcessCreationType string

const (
	Windows10EndpointProtectionConfigurationDefenderProcessCreationType_AuditMode   Windows10EndpointProtectionConfigurationDefenderProcessCreationType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Block       Windows10EndpointProtectionConfigurationDefenderProcessCreationType = "block"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Disable     Windows10EndpointProtectionConfigurationDefenderProcessCreationType = "disable"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationType_UserDefined Windows10EndpointProtectionConfigurationDefenderProcessCreationType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Warn        Windows10EndpointProtectionConfigurationDefenderProcessCreationType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderProcessCreationType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderProcessCreationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderProcessCreationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderProcessCreationType(input string) (*Windows10EndpointProtectionConfigurationDefenderProcessCreationType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderProcessCreationType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderProcessCreationType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderProcessCreationType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderProcessCreationType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderProcessCreationType(input)
	return &out, nil
}
