package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningSelfInitiatedCourseStatus string

const (
	LearningSelfInitiatedCourseStatus_Completed  LearningSelfInitiatedCourseStatus = "completed"
	LearningSelfInitiatedCourseStatus_InProgress LearningSelfInitiatedCourseStatus = "inProgress"
	LearningSelfInitiatedCourseStatus_NotStarted LearningSelfInitiatedCourseStatus = "notStarted"
)

func PossibleValuesForLearningSelfInitiatedCourseStatus() []string {
	return []string{
		string(LearningSelfInitiatedCourseStatus_Completed),
		string(LearningSelfInitiatedCourseStatus_InProgress),
		string(LearningSelfInitiatedCourseStatus_NotStarted),
	}
}

func (s *LearningSelfInitiatedCourseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLearningSelfInitiatedCourseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLearningSelfInitiatedCourseStatus(input string) (*LearningSelfInitiatedCourseStatus, error) {
	vals := map[string]LearningSelfInitiatedCourseStatus{
		"completed":  LearningSelfInitiatedCourseStatus_Completed,
		"inprogress": LearningSelfInitiatedCourseStatus_InProgress,
		"notstarted": LearningSelfInitiatedCourseStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningSelfInitiatedCourseStatus(input)
	return &out, nil
}
