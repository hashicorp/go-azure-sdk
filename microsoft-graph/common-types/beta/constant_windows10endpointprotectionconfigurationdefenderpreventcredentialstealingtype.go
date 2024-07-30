package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType string

const (
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_AuditMode     Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_Enable        Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "enable"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_NotConfigured Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_UserDefined   Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_Warn          Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType(input string) (*Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderPreventCredentialStealingType(input)
	return &out, nil
}
