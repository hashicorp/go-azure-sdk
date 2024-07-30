package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainerType string

const (
	PlannerPlanContainerType_Group  PlannerPlanContainerType = "group"
	PlannerPlanContainerType_Roster PlannerPlanContainerType = "roster"
)

func PossibleValuesForPlannerPlanContainerType() []string {
	return []string{
		string(PlannerPlanContainerType_Group),
		string(PlannerPlanContainerType_Roster),
	}
}

func (s *PlannerPlanContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanContainerType(input string) (*PlannerPlanContainerType, error) {
	vals := map[string]PlannerPlanContainerType{
		"group":  PlannerPlanContainerType_Group,
		"roster": PlannerPlanContainerType_Roster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContainerType(input)
	return &out, nil
}
