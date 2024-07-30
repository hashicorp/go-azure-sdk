package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderEmailContentExecution string

const (
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_AuditMode     Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_Enable        Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "enable"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_NotConfigured Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_UserDefined   Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_Warn          Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderEmailContentExecution() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderEmailContentExecution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderEmailContentExecution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderEmailContentExecution(input string) (*Windows10EndpointProtectionConfigurationDefenderEmailContentExecution, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderEmailContentExecution{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderEmailContentExecution_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderEmailContentExecution(input)
	return &out, nil
}
