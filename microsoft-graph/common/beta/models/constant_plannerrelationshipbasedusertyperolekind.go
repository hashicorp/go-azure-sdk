package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRelationshipBasedUserTypeRoleKind string

const (
	PlannerRelationshipBasedUserTypeRoleKindrelationship PlannerRelationshipBasedUserTypeRoleKind = "Relationship"
)

func PossibleValuesForPlannerRelationshipBasedUserTypeRoleKind() []string {
	return []string{
		string(PlannerRelationshipBasedUserTypeRoleKindrelationship),
	}
}

func parsePlannerRelationshipBasedUserTypeRoleKind(input string) (*PlannerRelationshipBasedUserTypeRoleKind, error) {
	vals := map[string]PlannerRelationshipBasedUserTypeRoleKind{
		"relationship": PlannerRelationshipBasedUserTypeRoleKindrelationship,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerRelationshipBasedUserTypeRoleKind(input)
	return &out, nil
}
