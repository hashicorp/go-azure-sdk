package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertRuleSeverity string

const (
	ManagedTenantsManagedTenantAlertRuleSeverity_High          ManagedTenantsManagedTenantAlertRuleSeverity = "high"
	ManagedTenantsManagedTenantAlertRuleSeverity_Informational ManagedTenantsManagedTenantAlertRuleSeverity = "informational"
	ManagedTenantsManagedTenantAlertRuleSeverity_Low           ManagedTenantsManagedTenantAlertRuleSeverity = "low"
	ManagedTenantsManagedTenantAlertRuleSeverity_Medium        ManagedTenantsManagedTenantAlertRuleSeverity = "medium"
	ManagedTenantsManagedTenantAlertRuleSeverity_Unknown       ManagedTenantsManagedTenantAlertRuleSeverity = "unknown"
)

func PossibleValuesForManagedTenantsManagedTenantAlertRuleSeverity() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertRuleSeverity_High),
		string(ManagedTenantsManagedTenantAlertRuleSeverity_Informational),
		string(ManagedTenantsManagedTenantAlertRuleSeverity_Low),
		string(ManagedTenantsManagedTenantAlertRuleSeverity_Medium),
		string(ManagedTenantsManagedTenantAlertRuleSeverity_Unknown),
	}
}

func (s *ManagedTenantsManagedTenantAlertRuleSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagedTenantAlertRuleSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagedTenantAlertRuleSeverity(input string) (*ManagedTenantsManagedTenantAlertRuleSeverity, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertRuleSeverity{
		"high":          ManagedTenantsManagedTenantAlertRuleSeverity_High,
		"informational": ManagedTenantsManagedTenantAlertRuleSeverity_Informational,
		"low":           ManagedTenantsManagedTenantAlertRuleSeverity_Low,
		"medium":        ManagedTenantsManagedTenantAlertRuleSeverity_Medium,
		"unknown":       ManagedTenantsManagedTenantAlertRuleSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertRuleSeverity(input)
	return &out, nil
}
