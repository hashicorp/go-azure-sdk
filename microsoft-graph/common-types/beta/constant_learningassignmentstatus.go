package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningAssignmentStatus string

const (
	LearningAssignmentStatus_Completed  LearningAssignmentStatus = "completed"
	LearningAssignmentStatus_InProgress LearningAssignmentStatus = "inProgress"
	LearningAssignmentStatus_NotStarted LearningAssignmentStatus = "notStarted"
)

func PossibleValuesForLearningAssignmentStatus() []string {
	return []string{
		string(LearningAssignmentStatus_Completed),
		string(LearningAssignmentStatus_InProgress),
		string(LearningAssignmentStatus_NotStarted),
	}
}

func (s *LearningAssignmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLearningAssignmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLearningAssignmentStatus(input string) (*LearningAssignmentStatus, error) {
	vals := map[string]LearningAssignmentStatus{
		"completed":  LearningAssignmentStatus_Completed,
		"inprogress": LearningAssignmentStatus_InProgress,
		"notstarted": LearningAssignmentStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningAssignmentStatus(input)
	return &out, nil
}
