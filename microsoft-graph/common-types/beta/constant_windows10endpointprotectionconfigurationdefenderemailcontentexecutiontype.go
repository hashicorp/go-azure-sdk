package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType string

const (
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_AuditMode   Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Block       Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "block"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Disable     Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "disable"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_UserDefined Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Warn        Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderEmailContentExecutionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderEmailContentExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderEmailContentExecutionType(input string) (*Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType(input)
	return &out, nil
}
