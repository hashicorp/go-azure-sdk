package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningAssignmentAssignmentType string

const (
	LearningAssignmentAssignmentType_Recommended LearningAssignmentAssignmentType = "recommended"
	LearningAssignmentAssignmentType_Required    LearningAssignmentAssignmentType = "required"
)

func PossibleValuesForLearningAssignmentAssignmentType() []string {
	return []string{
		string(LearningAssignmentAssignmentType_Recommended),
		string(LearningAssignmentAssignmentType_Required),
	}
}

func (s *LearningAssignmentAssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLearningAssignmentAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLearningAssignmentAssignmentType(input string) (*LearningAssignmentAssignmentType, error) {
	vals := map[string]LearningAssignmentAssignmentType{
		"recommended": LearningAssignmentAssignmentType_Recommended,
		"required":    LearningAssignmentAssignmentType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningAssignmentAssignmentType(input)
	return &out, nil
}
