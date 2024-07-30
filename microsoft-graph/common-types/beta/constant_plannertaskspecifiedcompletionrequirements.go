package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskSpecifiedCompletionRequirements string

const (
	PlannerTaskSpecifiedCompletionRequirements_ChecklistCompletion PlannerTaskSpecifiedCompletionRequirements = "checklistCompletion"
	PlannerTaskSpecifiedCompletionRequirements_None                PlannerTaskSpecifiedCompletionRequirements = "none"
)

func PossibleValuesForPlannerTaskSpecifiedCompletionRequirements() []string {
	return []string{
		string(PlannerTaskSpecifiedCompletionRequirements_ChecklistCompletion),
		string(PlannerTaskSpecifiedCompletionRequirements_None),
	}
}

func (s *PlannerTaskSpecifiedCompletionRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskSpecifiedCompletionRequirements(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskSpecifiedCompletionRequirements(input string) (*PlannerTaskSpecifiedCompletionRequirements, error) {
	vals := map[string]PlannerTaskSpecifiedCompletionRequirements{
		"checklistcompletion": PlannerTaskSpecifiedCompletionRequirements_ChecklistCompletion,
		"none":                PlannerTaskSpecifiedCompletionRequirements_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskSpecifiedCompletionRequirements(input)
	return &out, nil
}
