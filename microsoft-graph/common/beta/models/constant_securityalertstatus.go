package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertStatus string

const (
	SecurityAlertStatusinProgress SecurityAlertStatus = "InProgress"
	SecurityAlertStatusnew        SecurityAlertStatus = "New"
	SecurityAlertStatusresolved   SecurityAlertStatus = "Resolved"
	SecurityAlertStatusunknown    SecurityAlertStatus = "Unknown"
)

func PossibleValuesForSecurityAlertStatus() []string {
	return []string{
		string(SecurityAlertStatusinProgress),
		string(SecurityAlertStatusnew),
		string(SecurityAlertStatusresolved),
		string(SecurityAlertStatusunknown),
	}
}

func parseSecurityAlertStatus(input string) (*SecurityAlertStatus, error) {
	vals := map[string]SecurityAlertStatus{
		"inprogress": SecurityAlertStatusinProgress,
		"new":        SecurityAlertStatusnew,
		"resolved":   SecurityAlertStatusresolved,
		"unknown":    SecurityAlertStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertStatus(input)
	return &out, nil
}
