package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertSeverity string

const (
	ManagedTenantsManagedTenantAlertSeverity_High          ManagedTenantsManagedTenantAlertSeverity = "high"
	ManagedTenantsManagedTenantAlertSeverity_Informational ManagedTenantsManagedTenantAlertSeverity = "informational"
	ManagedTenantsManagedTenantAlertSeverity_Low           ManagedTenantsManagedTenantAlertSeverity = "low"
	ManagedTenantsManagedTenantAlertSeverity_Medium        ManagedTenantsManagedTenantAlertSeverity = "medium"
	ManagedTenantsManagedTenantAlertSeverity_Unknown       ManagedTenantsManagedTenantAlertSeverity = "unknown"
)

func PossibleValuesForManagedTenantsManagedTenantAlertSeverity() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertSeverity_High),
		string(ManagedTenantsManagedTenantAlertSeverity_Informational),
		string(ManagedTenantsManagedTenantAlertSeverity_Low),
		string(ManagedTenantsManagedTenantAlertSeverity_Medium),
		string(ManagedTenantsManagedTenantAlertSeverity_Unknown),
	}
}

func (s *ManagedTenantsManagedTenantAlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagedTenantAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagedTenantAlertSeverity(input string) (*ManagedTenantsManagedTenantAlertSeverity, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertSeverity{
		"high":          ManagedTenantsManagedTenantAlertSeverity_High,
		"informational": ManagedTenantsManagedTenantAlertSeverity_Informational,
		"low":           ManagedTenantsManagedTenantAlertSeverity_Low,
		"medium":        ManagedTenantsManagedTenantAlertSeverity_Medium,
		"unknown":       ManagedTenantsManagedTenantAlertSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertSeverity(input)
	return &out, nil
}
