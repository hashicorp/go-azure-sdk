package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryStateStatus string

const (
	AlertHistoryStateStatusdismissed  AlertHistoryStateStatus = "Dismissed"
	AlertHistoryStateStatusinProgress AlertHistoryStateStatus = "InProgress"
	AlertHistoryStateStatusnewAlert   AlertHistoryStateStatus = "NewAlert"
	AlertHistoryStateStatusresolved   AlertHistoryStateStatus = "Resolved"
	AlertHistoryStateStatusunknown    AlertHistoryStateStatus = "Unknown"
)

func PossibleValuesForAlertHistoryStateStatus() []string {
	return []string{
		string(AlertHistoryStateStatusdismissed),
		string(AlertHistoryStateStatusinProgress),
		string(AlertHistoryStateStatusnewAlert),
		string(AlertHistoryStateStatusresolved),
		string(AlertHistoryStateStatusunknown),
	}
}

func parseAlertHistoryStateStatus(input string) (*AlertHistoryStateStatus, error) {
	vals := map[string]AlertHistoryStateStatus{
		"dismissed":  AlertHistoryStateStatusdismissed,
		"inprogress": AlertHistoryStateStatusinProgress,
		"newalert":   AlertHistoryStateStatusnewAlert,
		"resolved":   AlertHistoryStateStatusresolved,
		"unknown":    AlertHistoryStateStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertHistoryStateStatus(input)
	return &out, nil
}
