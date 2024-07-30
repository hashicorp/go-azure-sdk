package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRelationshipBasedUserTypeRoleKind string

const (
	PlannerRelationshipBasedUserTypeRoleKind_Relationship PlannerRelationshipBasedUserTypeRoleKind = "relationship"
)

func PossibleValuesForPlannerRelationshipBasedUserTypeRoleKind() []string {
	return []string{
		string(PlannerRelationshipBasedUserTypeRoleKind_Relationship),
	}
}

func (s *PlannerRelationshipBasedUserTypeRoleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerRelationshipBasedUserTypeRoleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerRelationshipBasedUserTypeRoleKind(input string) (*PlannerRelationshipBasedUserTypeRoleKind, error) {
	vals := map[string]PlannerRelationshipBasedUserTypeRoleKind{
		"relationship": PlannerRelationshipBasedUserTypeRoleKind_Relationship,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerRelationshipBasedUserTypeRoleKind(input)
	return &out, nil
}
