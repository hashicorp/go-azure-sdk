package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRelationshipBasedUserTypeRole string

const (
	PlannerRelationshipBasedUserTypeRoleapplications  PlannerRelationshipBasedUserTypeRole = "Applications"
	PlannerRelationshipBasedUserTypeRoledefaultRules  PlannerRelationshipBasedUserTypeRole = "DefaultRules"
	PlannerRelationshipBasedUserTypeRolegroupMembers  PlannerRelationshipBasedUserTypeRole = "GroupMembers"
	PlannerRelationshipBasedUserTypeRolegroupOwners   PlannerRelationshipBasedUserTypeRole = "GroupOwners"
	PlannerRelationshipBasedUserTypeRoletaskAssignees PlannerRelationshipBasedUserTypeRole = "TaskAssignees"
)

func PossibleValuesForPlannerRelationshipBasedUserTypeRole() []string {
	return []string{
		string(PlannerRelationshipBasedUserTypeRoleapplications),
		string(PlannerRelationshipBasedUserTypeRoledefaultRules),
		string(PlannerRelationshipBasedUserTypeRolegroupMembers),
		string(PlannerRelationshipBasedUserTypeRolegroupOwners),
		string(PlannerRelationshipBasedUserTypeRoletaskAssignees),
	}
}

func parsePlannerRelationshipBasedUserTypeRole(input string) (*PlannerRelationshipBasedUserTypeRole, error) {
	vals := map[string]PlannerRelationshipBasedUserTypeRole{
		"applications":  PlannerRelationshipBasedUserTypeRoleapplications,
		"defaultrules":  PlannerRelationshipBasedUserTypeRoledefaultRules,
		"groupmembers":  PlannerRelationshipBasedUserTypeRolegroupMembers,
		"groupowners":   PlannerRelationshipBasedUserTypeRolegroupOwners,
		"taskassignees": PlannerRelationshipBasedUserTypeRoletaskAssignees,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerRelationshipBasedUserTypeRole(input)
	return &out, nil
}
