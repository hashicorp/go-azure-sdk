package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity string

const (
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_High          ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "high"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Informational ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "informational"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Low           ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "low"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Medium        ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "medium"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Unknown       ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "unknown"
)

func PossibleValuesForManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity() []string {
	return []string{
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_High),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Informational),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Low),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Medium),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Unknown),
	}
}

func (s *ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity(input string) (*ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity, error) {
	vals := map[string]ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity{
		"high":          ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_High,
		"informational": ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Informational,
		"low":           ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Low,
		"medium":        ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Medium,
		"unknown":       ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity(input)
	return &out, nil
}
