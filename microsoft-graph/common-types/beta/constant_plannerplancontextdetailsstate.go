package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContextDetailsState string

const (
	PlannerPlanContextDetailsState_Active   PlannerPlanContextDetailsState = "active"
	PlannerPlanContextDetailsState_Delinked PlannerPlanContextDetailsState = "delinked"
)

func PossibleValuesForPlannerPlanContextDetailsState() []string {
	return []string{
		string(PlannerPlanContextDetailsState_Active),
		string(PlannerPlanContextDetailsState_Delinked),
	}
}

func (s *PlannerPlanContextDetailsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanContextDetailsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanContextDetailsState(input string) (*PlannerPlanContextDetailsState, error) {
	vals := map[string]PlannerPlanContextDetailsState{
		"active":   PlannerPlanContextDetailsState_Active,
		"delinked": PlannerPlanContextDetailsState_Delinked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContextDetailsState(input)
	return &out, nil
}
