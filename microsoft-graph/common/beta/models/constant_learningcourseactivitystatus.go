package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningCourseActivityStatus string

const (
	LearningCourseActivityStatuscompleted  LearningCourseActivityStatus = "Completed"
	LearningCourseActivityStatusinProgress LearningCourseActivityStatus = "InProgress"
	LearningCourseActivityStatusnotStarted LearningCourseActivityStatus = "NotStarted"
)

func PossibleValuesForLearningCourseActivityStatus() []string {
	return []string{
		string(LearningCourseActivityStatuscompleted),
		string(LearningCourseActivityStatusinProgress),
		string(LearningCourseActivityStatusnotStarted),
	}
}

func parseLearningCourseActivityStatus(input string) (*LearningCourseActivityStatus, error) {
	vals := map[string]LearningCourseActivityStatus{
		"completed":  LearningCourseActivityStatuscompleted,
		"inprogress": LearningCourseActivityStatusinProgress,
		"notstarted": LearningCourseActivityStatusnotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningCourseActivityStatus(input)
	return &out, nil
}
