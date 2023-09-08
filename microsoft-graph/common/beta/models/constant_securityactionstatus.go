package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionStatus string

const (
	SecurityActionStatusCompleted  SecurityActionStatus = "Completed"
	SecurityActionStatusFailed     SecurityActionStatus = "Failed"
	SecurityActionStatusNotStarted SecurityActionStatus = "NotStarted"
	SecurityActionStatusRunning    SecurityActionStatus = "Running"
)

func PossibleValuesForSecurityActionStatus() []string {
	return []string{
		string(SecurityActionStatusCompleted),
		string(SecurityActionStatusFailed),
		string(SecurityActionStatusNotStarted),
		string(SecurityActionStatusRunning),
	}
}

func parseSecurityActionStatus(input string) (*SecurityActionStatus, error) {
	vals := map[string]SecurityActionStatus{
		"completed":  SecurityActionStatusCompleted,
		"failed":     SecurityActionStatusFailed,
		"notstarted": SecurityActionStatusNotStarted,
		"running":    SecurityActionStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityActionStatus(input)
	return &out, nil
}
