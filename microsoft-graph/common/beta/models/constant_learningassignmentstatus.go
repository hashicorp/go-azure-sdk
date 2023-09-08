package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningAssignmentStatus string

const (
	LearningAssignmentStatuscompleted  LearningAssignmentStatus = "Completed"
	LearningAssignmentStatusinProgress LearningAssignmentStatus = "InProgress"
	LearningAssignmentStatusnotStarted LearningAssignmentStatus = "NotStarted"
)

func PossibleValuesForLearningAssignmentStatus() []string {
	return []string{
		string(LearningAssignmentStatuscompleted),
		string(LearningAssignmentStatusinProgress),
		string(LearningAssignmentStatusnotStarted),
	}
}

func parseLearningAssignmentStatus(input string) (*LearningAssignmentStatus, error) {
	vals := map[string]LearningAssignmentStatus{
		"completed":  LearningAssignmentStatuscompleted,
		"inprogress": LearningAssignmentStatusinProgress,
		"notstarted": LearningAssignmentStatusnotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningAssignmentStatus(input)
	return &out, nil
}
