package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType string

const (
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_AuditMode     Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_Enable        Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "enable"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_NotConfigured Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_UserDefined   Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_Warn          Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderNetworkProtectionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderNetworkProtectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderNetworkProtectionType(input string) (*Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderNetworkProtectionType(input)
	return &out, nil
}
