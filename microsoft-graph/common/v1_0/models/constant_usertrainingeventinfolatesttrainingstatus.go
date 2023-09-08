package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingEventInfoLatestTrainingStatus string

const (
	UserTrainingEventInfoLatestTrainingStatusassigned   UserTrainingEventInfoLatestTrainingStatus = "Assigned"
	UserTrainingEventInfoLatestTrainingStatuscompleted  UserTrainingEventInfoLatestTrainingStatus = "Completed"
	UserTrainingEventInfoLatestTrainingStatusinProgress UserTrainingEventInfoLatestTrainingStatus = "InProgress"
	UserTrainingEventInfoLatestTrainingStatusoverdue    UserTrainingEventInfoLatestTrainingStatus = "Overdue"
	UserTrainingEventInfoLatestTrainingStatusunknown    UserTrainingEventInfoLatestTrainingStatus = "Unknown"
)

func PossibleValuesForUserTrainingEventInfoLatestTrainingStatus() []string {
	return []string{
		string(UserTrainingEventInfoLatestTrainingStatusassigned),
		string(UserTrainingEventInfoLatestTrainingStatuscompleted),
		string(UserTrainingEventInfoLatestTrainingStatusinProgress),
		string(UserTrainingEventInfoLatestTrainingStatusoverdue),
		string(UserTrainingEventInfoLatestTrainingStatusunknown),
	}
}

func parseUserTrainingEventInfoLatestTrainingStatus(input string) (*UserTrainingEventInfoLatestTrainingStatus, error) {
	vals := map[string]UserTrainingEventInfoLatestTrainingStatus{
		"assigned":   UserTrainingEventInfoLatestTrainingStatusassigned,
		"completed":  UserTrainingEventInfoLatestTrainingStatuscompleted,
		"inprogress": UserTrainingEventInfoLatestTrainingStatusinProgress,
		"overdue":    UserTrainingEventInfoLatestTrainingStatusoverdue,
		"unknown":    UserTrainingEventInfoLatestTrainingStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserTrainingEventInfoLatestTrainingStatus(input)
	return &out, nil
}
