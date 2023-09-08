package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingStatusInfoTrainingStatus string

const (
	UserTrainingStatusInfoTrainingStatusassigned   UserTrainingStatusInfoTrainingStatus = "Assigned"
	UserTrainingStatusInfoTrainingStatuscompleted  UserTrainingStatusInfoTrainingStatus = "Completed"
	UserTrainingStatusInfoTrainingStatusinProgress UserTrainingStatusInfoTrainingStatus = "InProgress"
	UserTrainingStatusInfoTrainingStatusoverdue    UserTrainingStatusInfoTrainingStatus = "Overdue"
	UserTrainingStatusInfoTrainingStatusunknown    UserTrainingStatusInfoTrainingStatus = "Unknown"
)

func PossibleValuesForUserTrainingStatusInfoTrainingStatus() []string {
	return []string{
		string(UserTrainingStatusInfoTrainingStatusassigned),
		string(UserTrainingStatusInfoTrainingStatuscompleted),
		string(UserTrainingStatusInfoTrainingStatusinProgress),
		string(UserTrainingStatusInfoTrainingStatusoverdue),
		string(UserTrainingStatusInfoTrainingStatusunknown),
	}
}

func parseUserTrainingStatusInfoTrainingStatus(input string) (*UserTrainingStatusInfoTrainingStatus, error) {
	vals := map[string]UserTrainingStatusInfoTrainingStatus{
		"assigned":   UserTrainingStatusInfoTrainingStatusassigned,
		"completed":  UserTrainingStatusInfoTrainingStatuscompleted,
		"inprogress": UserTrainingStatusInfoTrainingStatusinProgress,
		"overdue":    UserTrainingStatusInfoTrainingStatusoverdue,
		"unknown":    UserTrainingStatusInfoTrainingStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserTrainingStatusInfoTrainingStatus(input)
	return &out, nil
}
