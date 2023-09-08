package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterTypeAndEvaluationResultEvaluationResult string

const (
	AssignmentFilterTypeAndEvaluationResultEvaluationResultfailure      AssignmentFilterTypeAndEvaluationResultEvaluationResult = "Failure"
	AssignmentFilterTypeAndEvaluationResultEvaluationResultinconclusive AssignmentFilterTypeAndEvaluationResultEvaluationResult = "Inconclusive"
	AssignmentFilterTypeAndEvaluationResultEvaluationResultmatch        AssignmentFilterTypeAndEvaluationResultEvaluationResult = "Match"
	AssignmentFilterTypeAndEvaluationResultEvaluationResultnotEvaluated AssignmentFilterTypeAndEvaluationResultEvaluationResult = "NotEvaluated"
	AssignmentFilterTypeAndEvaluationResultEvaluationResultnotMatch     AssignmentFilterTypeAndEvaluationResultEvaluationResult = "NotMatch"
	AssignmentFilterTypeAndEvaluationResultEvaluationResultunknown      AssignmentFilterTypeAndEvaluationResultEvaluationResult = "Unknown"
)

func PossibleValuesForAssignmentFilterTypeAndEvaluationResultEvaluationResult() []string {
	return []string{
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResultfailure),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResultinconclusive),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResultmatch),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResultnotEvaluated),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResultnotMatch),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResultunknown),
	}
}

func parseAssignmentFilterTypeAndEvaluationResultEvaluationResult(input string) (*AssignmentFilterTypeAndEvaluationResultEvaluationResult, error) {
	vals := map[string]AssignmentFilterTypeAndEvaluationResultEvaluationResult{
		"failure":      AssignmentFilterTypeAndEvaluationResultEvaluationResultfailure,
		"inconclusive": AssignmentFilterTypeAndEvaluationResultEvaluationResultinconclusive,
		"match":        AssignmentFilterTypeAndEvaluationResultEvaluationResultmatch,
		"notevaluated": AssignmentFilterTypeAndEvaluationResultEvaluationResultnotEvaluated,
		"notmatch":     AssignmentFilterTypeAndEvaluationResultEvaluationResultnotMatch,
		"unknown":      AssignmentFilterTypeAndEvaluationResultEvaluationResultunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterTypeAndEvaluationResultEvaluationResult(input)
	return &out, nil
}
