package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType string

const (
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_AuditMode     Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_Enable        Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "enable"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_NotConfigured Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_UserDefined   Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_Warn          Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType(input string) (*Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderAdvancedRansomewareProtectionType(input)
	return &out, nil
}
