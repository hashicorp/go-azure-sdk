package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertStatus string

const (
	AlertStatusdismissed  AlertStatus = "Dismissed"
	AlertStatusinProgress AlertStatus = "InProgress"
	AlertStatusnewAlert   AlertStatus = "NewAlert"
	AlertStatusresolved   AlertStatus = "Resolved"
	AlertStatusunknown    AlertStatus = "Unknown"
)

func PossibleValuesForAlertStatus() []string {
	return []string{
		string(AlertStatusdismissed),
		string(AlertStatusinProgress),
		string(AlertStatusnewAlert),
		string(AlertStatusresolved),
		string(AlertStatusunknown),
	}
}

func parseAlertStatus(input string) (*AlertStatus, error) {
	vals := map[string]AlertStatus{
		"dismissed":  AlertStatusdismissed,
		"inprogress": AlertStatusinProgress,
		"newalert":   AlertStatusnewAlert,
		"resolved":   AlertStatusresolved,
		"unknown":    AlertStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertStatus(input)
	return &out, nil
}
