package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskStatus string

const (
	OutlookTaskStatuscompleted       OutlookTaskStatus = "Completed"
	OutlookTaskStatusdeferred        OutlookTaskStatus = "Deferred"
	OutlookTaskStatusinProgress      OutlookTaskStatus = "InProgress"
	OutlookTaskStatusnotStarted      OutlookTaskStatus = "NotStarted"
	OutlookTaskStatuswaitingOnOthers OutlookTaskStatus = "WaitingOnOthers"
)

func PossibleValuesForOutlookTaskStatus() []string {
	return []string{
		string(OutlookTaskStatuscompleted),
		string(OutlookTaskStatusdeferred),
		string(OutlookTaskStatusinProgress),
		string(OutlookTaskStatusnotStarted),
		string(OutlookTaskStatuswaitingOnOthers),
	}
}

func parseOutlookTaskStatus(input string) (*OutlookTaskStatus, error) {
	vals := map[string]OutlookTaskStatus{
		"completed":       OutlookTaskStatuscompleted,
		"deferred":        OutlookTaskStatusdeferred,
		"inprogress":      OutlookTaskStatusinProgress,
		"notstarted":      OutlookTaskStatusnotStarted,
		"waitingonothers": OutlookTaskStatuswaitingOnOthers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskStatus(input)
	return &out, nil
}
