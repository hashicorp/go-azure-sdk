package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskConfigurationRoleBaseRoleKind string

const (
	PlannerTaskConfigurationRoleBaseRoleKind_Relationship PlannerTaskConfigurationRoleBaseRoleKind = "relationship"
)

func PossibleValuesForPlannerTaskConfigurationRoleBaseRoleKind() []string {
	return []string{
		string(PlannerTaskConfigurationRoleBaseRoleKind_Relationship),
	}
}

func (s *PlannerTaskConfigurationRoleBaseRoleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskConfigurationRoleBaseRoleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskConfigurationRoleBaseRoleKind(input string) (*PlannerTaskConfigurationRoleBaseRoleKind, error) {
	vals := map[string]PlannerTaskConfigurationRoleBaseRoleKind{
		"relationship": PlannerTaskConfigurationRoleBaseRoleKind_Relationship,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskConfigurationRoleBaseRoleKind(input)
	return &out, nil
}
