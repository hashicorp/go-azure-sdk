package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterTypeAndEvaluationResultAssignmentFilterType string

const (
	AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_Exclude AssignmentFilterTypeAndEvaluationResultAssignmentFilterType = "exclude"
	AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_Include AssignmentFilterTypeAndEvaluationResultAssignmentFilterType = "include"
	AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_None    AssignmentFilterTypeAndEvaluationResultAssignmentFilterType = "none"
)

func PossibleValuesForAssignmentFilterTypeAndEvaluationResultAssignmentFilterType() []string {
	return []string{
		string(AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_Exclude),
		string(AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_Include),
		string(AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_None),
	}
}

func (s *AssignmentFilterTypeAndEvaluationResultAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterTypeAndEvaluationResultAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterTypeAndEvaluationResultAssignmentFilterType(input string) (*AssignmentFilterTypeAndEvaluationResultAssignmentFilterType, error) {
	vals := map[string]AssignmentFilterTypeAndEvaluationResultAssignmentFilterType{
		"exclude": AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_Exclude,
		"include": AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_Include,
		"none":    AssignmentFilterTypeAndEvaluationResultAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterTypeAndEvaluationResultAssignmentFilterType(input)
	return &out, nil
}
