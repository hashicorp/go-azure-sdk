package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRelationshipBasedUserTypeRole string

const (
	PlannerRelationshipBasedUserTypeRole_Applications  PlannerRelationshipBasedUserTypeRole = "applications"
	PlannerRelationshipBasedUserTypeRole_DefaultRules  PlannerRelationshipBasedUserTypeRole = "defaultRules"
	PlannerRelationshipBasedUserTypeRole_GroupMembers  PlannerRelationshipBasedUserTypeRole = "groupMembers"
	PlannerRelationshipBasedUserTypeRole_GroupOwners   PlannerRelationshipBasedUserTypeRole = "groupOwners"
	PlannerRelationshipBasedUserTypeRole_TaskAssignees PlannerRelationshipBasedUserTypeRole = "taskAssignees"
)

func PossibleValuesForPlannerRelationshipBasedUserTypeRole() []string {
	return []string{
		string(PlannerRelationshipBasedUserTypeRole_Applications),
		string(PlannerRelationshipBasedUserTypeRole_DefaultRules),
		string(PlannerRelationshipBasedUserTypeRole_GroupMembers),
		string(PlannerRelationshipBasedUserTypeRole_GroupOwners),
		string(PlannerRelationshipBasedUserTypeRole_TaskAssignees),
	}
}

func (s *PlannerRelationshipBasedUserTypeRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerRelationshipBasedUserTypeRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerRelationshipBasedUserTypeRole(input string) (*PlannerRelationshipBasedUserTypeRole, error) {
	vals := map[string]PlannerRelationshipBasedUserTypeRole{
		"applications":  PlannerRelationshipBasedUserTypeRole_Applications,
		"defaultrules":  PlannerRelationshipBasedUserTypeRole_DefaultRules,
		"groupmembers":  PlannerRelationshipBasedUserTypeRole_GroupMembers,
		"groupowners":   PlannerRelationshipBasedUserTypeRole_GroupOwners,
		"taskassignees": PlannerRelationshipBasedUserTypeRole_TaskAssignees,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerRelationshipBasedUserTypeRole(input)
	return &out, nil
}
