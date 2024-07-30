package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType string

const (
	Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_AuditMode   Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Block       Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType = "block"
	Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Disable     Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType = "disable"
	Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_UserDefined Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Warn        Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType(input string) (*Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderBlockPersistenceThroughWmiType(input)
	return &out, nil
}
