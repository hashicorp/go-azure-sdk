package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummaryAssignmentFilterType string

const (
	AssignmentFilterEvaluationSummaryAssignmentFilterTypeexclude AssignmentFilterEvaluationSummaryAssignmentFilterType = "Exclude"
	AssignmentFilterEvaluationSummaryAssignmentFilterTypeinclude AssignmentFilterEvaluationSummaryAssignmentFilterType = "Include"
	AssignmentFilterEvaluationSummaryAssignmentFilterTypenone    AssignmentFilterEvaluationSummaryAssignmentFilterType = "None"
)

func PossibleValuesForAssignmentFilterEvaluationSummaryAssignmentFilterType() []string {
	return []string{
		string(AssignmentFilterEvaluationSummaryAssignmentFilterTypeexclude),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterTypeinclude),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterTypenone),
	}
}

func parseAssignmentFilterEvaluationSummaryAssignmentFilterType(input string) (*AssignmentFilterEvaluationSummaryAssignmentFilterType, error) {
	vals := map[string]AssignmentFilterEvaluationSummaryAssignmentFilterType{
		"exclude": AssignmentFilterEvaluationSummaryAssignmentFilterTypeexclude,
		"include": AssignmentFilterEvaluationSummaryAssignmentFilterTypeinclude,
		"none":    AssignmentFilterEvaluationSummaryAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluationSummaryAssignmentFilterType(input)
	return &out, nil
}
