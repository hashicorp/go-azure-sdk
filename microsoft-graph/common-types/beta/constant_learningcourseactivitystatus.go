package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningCourseActivityStatus string

const (
	LearningCourseActivityStatus_Completed  LearningCourseActivityStatus = "completed"
	LearningCourseActivityStatus_InProgress LearningCourseActivityStatus = "inProgress"
	LearningCourseActivityStatus_NotStarted LearningCourseActivityStatus = "notStarted"
)

func PossibleValuesForLearningCourseActivityStatus() []string {
	return []string{
		string(LearningCourseActivityStatus_Completed),
		string(LearningCourseActivityStatus_InProgress),
		string(LearningCourseActivityStatus_NotStarted),
	}
}

func (s *LearningCourseActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLearningCourseActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLearningCourseActivityStatus(input string) (*LearningCourseActivityStatus, error) {
	vals := map[string]LearningCourseActivityStatus{
		"completed":  LearningCourseActivityStatus_Completed,
		"inprogress": LearningCourseActivityStatus_InProgress,
		"notstarted": LearningCourseActivityStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningCourseActivityStatus(input)
	return &out, nil
}
