package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskSpecifiedCompletionRequirements string

const (
	BusinessScenarioTaskSpecifiedCompletionRequirementschecklistCompletion BusinessScenarioTaskSpecifiedCompletionRequirements = "ChecklistCompletion"
	BusinessScenarioTaskSpecifiedCompletionRequirementsnone                BusinessScenarioTaskSpecifiedCompletionRequirements = "None"
)

func PossibleValuesForBusinessScenarioTaskSpecifiedCompletionRequirements() []string {
	return []string{
		string(BusinessScenarioTaskSpecifiedCompletionRequirementschecklistCompletion),
		string(BusinessScenarioTaskSpecifiedCompletionRequirementsnone),
	}
}

func parseBusinessScenarioTaskSpecifiedCompletionRequirements(input string) (*BusinessScenarioTaskSpecifiedCompletionRequirements, error) {
	vals := map[string]BusinessScenarioTaskSpecifiedCompletionRequirements{
		"checklistcompletion": BusinessScenarioTaskSpecifiedCompletionRequirementschecklistCompletion,
		"none":                BusinessScenarioTaskSpecifiedCompletionRequirementsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioTaskSpecifiedCompletionRequirements(input)
	return &out, nil
}
