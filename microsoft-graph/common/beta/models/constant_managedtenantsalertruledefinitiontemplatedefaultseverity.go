package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity string

const (
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityhigh          ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "High"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityinformational ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "Informational"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeveritylow           ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "Low"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeveritymedium        ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "Medium"
	ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityunknown       ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity = "Unknown"
)

func PossibleValuesForManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity() []string {
	return []string{
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityhigh),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityinformational),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeveritylow),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeveritymedium),
		string(ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityunknown),
	}
}

func parseManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity(input string) (*ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity, error) {
	vals := map[string]ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity{
		"high":          ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityhigh,
		"informational": ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityinformational,
		"low":           ManagedTenantsAlertRuleDefinitionTemplateDefaultSeveritylow,
		"medium":        ManagedTenantsAlertRuleDefinitionTemplateDefaultSeveritymedium,
		"unknown":       ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity(input)
	return &out, nil
}
