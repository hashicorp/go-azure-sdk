package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummaryAssignmentFilterType string

const (
	AssignmentFilterEvaluationSummaryAssignmentFilterType_Exclude AssignmentFilterEvaluationSummaryAssignmentFilterType = "exclude"
	AssignmentFilterEvaluationSummaryAssignmentFilterType_Include AssignmentFilterEvaluationSummaryAssignmentFilterType = "include"
	AssignmentFilterEvaluationSummaryAssignmentFilterType_None    AssignmentFilterEvaluationSummaryAssignmentFilterType = "none"
)

func PossibleValuesForAssignmentFilterEvaluationSummaryAssignmentFilterType() []string {
	return []string{
		string(AssignmentFilterEvaluationSummaryAssignmentFilterType_Exclude),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterType_Include),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterType_None),
	}
}

func (s *AssignmentFilterEvaluationSummaryAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterEvaluationSummaryAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterEvaluationSummaryAssignmentFilterType(input string) (*AssignmentFilterEvaluationSummaryAssignmentFilterType, error) {
	vals := map[string]AssignmentFilterEvaluationSummaryAssignmentFilterType{
		"exclude": AssignmentFilterEvaluationSummaryAssignmentFilterType_Exclude,
		"include": AssignmentFilterEvaluationSummaryAssignmentFilterType_Include,
		"none":    AssignmentFilterEvaluationSummaryAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluationSummaryAssignmentFilterType(input)
	return &out, nil
}
