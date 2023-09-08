package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummaryEvaluationResult string

const (
	AssignmentFilterEvaluationSummaryEvaluationResultfailure      AssignmentFilterEvaluationSummaryEvaluationResult = "Failure"
	AssignmentFilterEvaluationSummaryEvaluationResultinconclusive AssignmentFilterEvaluationSummaryEvaluationResult = "Inconclusive"
	AssignmentFilterEvaluationSummaryEvaluationResultmatch        AssignmentFilterEvaluationSummaryEvaluationResult = "Match"
	AssignmentFilterEvaluationSummaryEvaluationResultnotEvaluated AssignmentFilterEvaluationSummaryEvaluationResult = "NotEvaluated"
	AssignmentFilterEvaluationSummaryEvaluationResultnotMatch     AssignmentFilterEvaluationSummaryEvaluationResult = "NotMatch"
	AssignmentFilterEvaluationSummaryEvaluationResultunknown      AssignmentFilterEvaluationSummaryEvaluationResult = "Unknown"
)

func PossibleValuesForAssignmentFilterEvaluationSummaryEvaluationResult() []string {
	return []string{
		string(AssignmentFilterEvaluationSummaryEvaluationResultfailure),
		string(AssignmentFilterEvaluationSummaryEvaluationResultinconclusive),
		string(AssignmentFilterEvaluationSummaryEvaluationResultmatch),
		string(AssignmentFilterEvaluationSummaryEvaluationResultnotEvaluated),
		string(AssignmentFilterEvaluationSummaryEvaluationResultnotMatch),
		string(AssignmentFilterEvaluationSummaryEvaluationResultunknown),
	}
}

func parseAssignmentFilterEvaluationSummaryEvaluationResult(input string) (*AssignmentFilterEvaluationSummaryEvaluationResult, error) {
	vals := map[string]AssignmentFilterEvaluationSummaryEvaluationResult{
		"failure":      AssignmentFilterEvaluationSummaryEvaluationResultfailure,
		"inconclusive": AssignmentFilterEvaluationSummaryEvaluationResultinconclusive,
		"match":        AssignmentFilterEvaluationSummaryEvaluationResultmatch,
		"notevaluated": AssignmentFilterEvaluationSummaryEvaluationResultnotEvaluated,
		"notmatch":     AssignmentFilterEvaluationSummaryEvaluationResultnotMatch,
		"unknown":      AssignmentFilterEvaluationSummaryEvaluationResultunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluationSummaryEvaluationResult(input)
	return &out, nil
}
