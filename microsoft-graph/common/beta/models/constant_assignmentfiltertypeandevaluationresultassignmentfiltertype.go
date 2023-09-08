package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterTypeAndEvaluationResultAssignmentFilterType string

const (
	AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypeexclude AssignmentFilterTypeAndEvaluationResultAssignmentFilterType = "Exclude"
	AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypeinclude AssignmentFilterTypeAndEvaluationResultAssignmentFilterType = "Include"
	AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypenone    AssignmentFilterTypeAndEvaluationResultAssignmentFilterType = "None"
)

func PossibleValuesForAssignmentFilterTypeAndEvaluationResultAssignmentFilterType() []string {
	return []string{
		string(AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypeexclude),
		string(AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypeinclude),
		string(AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypenone),
	}
}

func parseAssignmentFilterTypeAndEvaluationResultAssignmentFilterType(input string) (*AssignmentFilterTypeAndEvaluationResultAssignmentFilterType, error) {
	vals := map[string]AssignmentFilterTypeAndEvaluationResultAssignmentFilterType{
		"exclude": AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypeexclude,
		"include": AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypeinclude,
		"none":    AssignmentFilterTypeAndEvaluationResultAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterTypeAndEvaluationResultAssignmentFilterType(input)
	return &out, nil
}
