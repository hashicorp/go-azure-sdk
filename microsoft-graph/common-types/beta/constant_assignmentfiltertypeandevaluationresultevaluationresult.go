package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterTypeAndEvaluationResultEvaluationResult string

const (
	AssignmentFilterTypeAndEvaluationResultEvaluationResult_Failure      AssignmentFilterTypeAndEvaluationResultEvaluationResult = "failure"
	AssignmentFilterTypeAndEvaluationResultEvaluationResult_Inconclusive AssignmentFilterTypeAndEvaluationResultEvaluationResult = "inconclusive"
	AssignmentFilterTypeAndEvaluationResultEvaluationResult_Match        AssignmentFilterTypeAndEvaluationResultEvaluationResult = "match"
	AssignmentFilterTypeAndEvaluationResultEvaluationResult_NotEvaluated AssignmentFilterTypeAndEvaluationResultEvaluationResult = "notEvaluated"
	AssignmentFilterTypeAndEvaluationResultEvaluationResult_NotMatch     AssignmentFilterTypeAndEvaluationResultEvaluationResult = "notMatch"
	AssignmentFilterTypeAndEvaluationResultEvaluationResult_Unknown      AssignmentFilterTypeAndEvaluationResultEvaluationResult = "unknown"
)

func PossibleValuesForAssignmentFilterTypeAndEvaluationResultEvaluationResult() []string {
	return []string{
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResult_Failure),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResult_Inconclusive),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResult_Match),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResult_NotEvaluated),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResult_NotMatch),
		string(AssignmentFilterTypeAndEvaluationResultEvaluationResult_Unknown),
	}
}

func (s *AssignmentFilterTypeAndEvaluationResultEvaluationResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterTypeAndEvaluationResultEvaluationResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterTypeAndEvaluationResultEvaluationResult(input string) (*AssignmentFilterTypeAndEvaluationResultEvaluationResult, error) {
	vals := map[string]AssignmentFilterTypeAndEvaluationResultEvaluationResult{
		"failure":      AssignmentFilterTypeAndEvaluationResultEvaluationResult_Failure,
		"inconclusive": AssignmentFilterTypeAndEvaluationResultEvaluationResult_Inconclusive,
		"match":        AssignmentFilterTypeAndEvaluationResultEvaluationResult_Match,
		"notevaluated": AssignmentFilterTypeAndEvaluationResultEvaluationResult_NotEvaluated,
		"notmatch":     AssignmentFilterTypeAndEvaluationResultEvaluationResult_NotMatch,
		"unknown":      AssignmentFilterTypeAndEvaluationResultEvaluationResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterTypeAndEvaluationResultEvaluationResult(input)
	return &out, nil
}
