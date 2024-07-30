package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskStatus string

const (
	SecurityConfigurationTaskStatus_Active    SecurityConfigurationTaskStatus = "active"
	SecurityConfigurationTaskStatus_Completed SecurityConfigurationTaskStatus = "completed"
	SecurityConfigurationTaskStatus_Pending   SecurityConfigurationTaskStatus = "pending"
	SecurityConfigurationTaskStatus_Rejected  SecurityConfigurationTaskStatus = "rejected"
	SecurityConfigurationTaskStatus_Unknown   SecurityConfigurationTaskStatus = "unknown"
)

func PossibleValuesForSecurityConfigurationTaskStatus() []string {
	return []string{
		string(SecurityConfigurationTaskStatus_Active),
		string(SecurityConfigurationTaskStatus_Completed),
		string(SecurityConfigurationTaskStatus_Pending),
		string(SecurityConfigurationTaskStatus_Rejected),
		string(SecurityConfigurationTaskStatus_Unknown),
	}
}

func (s *SecurityConfigurationTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityConfigurationTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityConfigurationTaskStatus(input string) (*SecurityConfigurationTaskStatus, error) {
	vals := map[string]SecurityConfigurationTaskStatus{
		"active":    SecurityConfigurationTaskStatus_Active,
		"completed": SecurityConfigurationTaskStatus_Completed,
		"pending":   SecurityConfigurationTaskStatus_Pending,
		"rejected":  SecurityConfigurationTaskStatus_Rejected,
		"unknown":   SecurityConfigurationTaskStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskStatus(input)
	return &out, nil
}
