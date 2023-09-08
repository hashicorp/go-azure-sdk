package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementAlertDefinitionSeverityLevel string

const (
	UnifiedRoleManagementAlertDefinitionSeverityLevelhigh          UnifiedRoleManagementAlertDefinitionSeverityLevel = "High"
	UnifiedRoleManagementAlertDefinitionSeverityLevelinformational UnifiedRoleManagementAlertDefinitionSeverityLevel = "Informational"
	UnifiedRoleManagementAlertDefinitionSeverityLevellow           UnifiedRoleManagementAlertDefinitionSeverityLevel = "Low"
	UnifiedRoleManagementAlertDefinitionSeverityLevelmedium        UnifiedRoleManagementAlertDefinitionSeverityLevel = "Medium"
	UnifiedRoleManagementAlertDefinitionSeverityLevelunknown       UnifiedRoleManagementAlertDefinitionSeverityLevel = "Unknown"
)

func PossibleValuesForUnifiedRoleManagementAlertDefinitionSeverityLevel() []string {
	return []string{
		string(UnifiedRoleManagementAlertDefinitionSeverityLevelhigh),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevelinformational),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevellow),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevelmedium),
		string(UnifiedRoleManagementAlertDefinitionSeverityLevelunknown),
	}
}

func parseUnifiedRoleManagementAlertDefinitionSeverityLevel(input string) (*UnifiedRoleManagementAlertDefinitionSeverityLevel, error) {
	vals := map[string]UnifiedRoleManagementAlertDefinitionSeverityLevel{
		"high":          UnifiedRoleManagementAlertDefinitionSeverityLevelhigh,
		"informational": UnifiedRoleManagementAlertDefinitionSeverityLevelinformational,
		"low":           UnifiedRoleManagementAlertDefinitionSeverityLevellow,
		"medium":        UnifiedRoleManagementAlertDefinitionSeverityLevelmedium,
		"unknown":       UnifiedRoleManagementAlertDefinitionSeverityLevelunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleManagementAlertDefinitionSeverityLevel(input)
	return &out, nil
}
