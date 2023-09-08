package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskConfigurationRoleBaseRoleKind string

const (
	PlannerTaskConfigurationRoleBaseRoleKindrelationship PlannerTaskConfigurationRoleBaseRoleKind = "Relationship"
)

func PossibleValuesForPlannerTaskConfigurationRoleBaseRoleKind() []string {
	return []string{
		string(PlannerTaskConfigurationRoleBaseRoleKindrelationship),
	}
}

func parsePlannerTaskConfigurationRoleBaseRoleKind(input string) (*PlannerTaskConfigurationRoleBaseRoleKind, error) {
	vals := map[string]PlannerTaskConfigurationRoleBaseRoleKind{
		"relationship": PlannerTaskConfigurationRoleBaseRoleKindrelationship,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskConfigurationRoleBaseRoleKind(input)
	return &out, nil
}
