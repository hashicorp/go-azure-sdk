package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskCompletionRequirements string

const (
	PlannerTaskCompletionRequirements_ChecklistCompletion PlannerTaskCompletionRequirements = "checklistCompletion"
	PlannerTaskCompletionRequirements_None                PlannerTaskCompletionRequirements = "none"
)

func PossibleValuesForPlannerTaskCompletionRequirements() []string {
	return []string{
		string(PlannerTaskCompletionRequirements_ChecklistCompletion),
		string(PlannerTaskCompletionRequirements_None),
	}
}

func (s *PlannerTaskCompletionRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskCompletionRequirements(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskCompletionRequirements(input string) (*PlannerTaskCompletionRequirements, error) {
	vals := map[string]PlannerTaskCompletionRequirements{
		"checklistcompletion": PlannerTaskCompletionRequirements_ChecklistCompletion,
		"none":                PlannerTaskCompletionRequirements_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskCompletionRequirements(input)
	return &out, nil
}
