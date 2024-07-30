package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummaryEvaluationResult string

const (
	AssignmentFilterEvaluationSummaryEvaluationResult_Failure      AssignmentFilterEvaluationSummaryEvaluationResult = "failure"
	AssignmentFilterEvaluationSummaryEvaluationResult_Inconclusive AssignmentFilterEvaluationSummaryEvaluationResult = "inconclusive"
	AssignmentFilterEvaluationSummaryEvaluationResult_Match        AssignmentFilterEvaluationSummaryEvaluationResult = "match"
	AssignmentFilterEvaluationSummaryEvaluationResult_NotEvaluated AssignmentFilterEvaluationSummaryEvaluationResult = "notEvaluated"
	AssignmentFilterEvaluationSummaryEvaluationResult_NotMatch     AssignmentFilterEvaluationSummaryEvaluationResult = "notMatch"
	AssignmentFilterEvaluationSummaryEvaluationResult_Unknown      AssignmentFilterEvaluationSummaryEvaluationResult = "unknown"
)

func PossibleValuesForAssignmentFilterEvaluationSummaryEvaluationResult() []string {
	return []string{
		string(AssignmentFilterEvaluationSummaryEvaluationResult_Failure),
		string(AssignmentFilterEvaluationSummaryEvaluationResult_Inconclusive),
		string(AssignmentFilterEvaluationSummaryEvaluationResult_Match),
		string(AssignmentFilterEvaluationSummaryEvaluationResult_NotEvaluated),
		string(AssignmentFilterEvaluationSummaryEvaluationResult_NotMatch),
		string(AssignmentFilterEvaluationSummaryEvaluationResult_Unknown),
	}
}

func (s *AssignmentFilterEvaluationSummaryEvaluationResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterEvaluationSummaryEvaluationResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterEvaluationSummaryEvaluationResult(input string) (*AssignmentFilterEvaluationSummaryEvaluationResult, error) {
	vals := map[string]AssignmentFilterEvaluationSummaryEvaluationResult{
		"failure":      AssignmentFilterEvaluationSummaryEvaluationResult_Failure,
		"inconclusive": AssignmentFilterEvaluationSummaryEvaluationResult_Inconclusive,
		"match":        AssignmentFilterEvaluationSummaryEvaluationResult_Match,
		"notevaluated": AssignmentFilterEvaluationSummaryEvaluationResult_NotEvaluated,
		"notmatch":     AssignmentFilterEvaluationSummaryEvaluationResult_NotMatch,
		"unknown":      AssignmentFilterEvaluationSummaryEvaluationResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluationSummaryEvaluationResult(input)
	return &out, nil
}
