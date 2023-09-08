package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskSpecifiedCompletionRequirements string

const (
	PlannerTaskSpecifiedCompletionRequirementschecklistCompletion PlannerTaskSpecifiedCompletionRequirements = "ChecklistCompletion"
	PlannerTaskSpecifiedCompletionRequirementsnone                PlannerTaskSpecifiedCompletionRequirements = "None"
)

func PossibleValuesForPlannerTaskSpecifiedCompletionRequirements() []string {
	return []string{
		string(PlannerTaskSpecifiedCompletionRequirementschecklistCompletion),
		string(PlannerTaskSpecifiedCompletionRequirementsnone),
	}
}

func parsePlannerTaskSpecifiedCompletionRequirements(input string) (*PlannerTaskSpecifiedCompletionRequirements, error) {
	vals := map[string]PlannerTaskSpecifiedCompletionRequirements{
		"checklistcompletion": PlannerTaskSpecifiedCompletionRequirementschecklistCompletion,
		"none":                PlannerTaskSpecifiedCompletionRequirementsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskSpecifiedCompletionRequirements(input)
	return &out, nil
}
