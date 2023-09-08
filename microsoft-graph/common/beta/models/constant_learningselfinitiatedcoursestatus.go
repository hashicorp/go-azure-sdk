package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningSelfInitiatedCourseStatus string

const (
	LearningSelfInitiatedCourseStatuscompleted  LearningSelfInitiatedCourseStatus = "Completed"
	LearningSelfInitiatedCourseStatusinProgress LearningSelfInitiatedCourseStatus = "InProgress"
	LearningSelfInitiatedCourseStatusnotStarted LearningSelfInitiatedCourseStatus = "NotStarted"
)

func PossibleValuesForLearningSelfInitiatedCourseStatus() []string {
	return []string{
		string(LearningSelfInitiatedCourseStatuscompleted),
		string(LearningSelfInitiatedCourseStatusinProgress),
		string(LearningSelfInitiatedCourseStatusnotStarted),
	}
}

func parseLearningSelfInitiatedCourseStatus(input string) (*LearningSelfInitiatedCourseStatus, error) {
	vals := map[string]LearningSelfInitiatedCourseStatus{
		"completed":  LearningSelfInitiatedCourseStatuscompleted,
		"inprogress": LearningSelfInitiatedCourseStatusinProgress,
		"notstarted": LearningSelfInitiatedCourseStatusnotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningSelfInitiatedCourseStatus(input)
	return &out, nil
}
