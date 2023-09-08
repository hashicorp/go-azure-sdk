package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionStateStatus string

const (
	SecurityActionStateStatusCompleted  SecurityActionStateStatus = "Completed"
	SecurityActionStateStatusFailed     SecurityActionStateStatus = "Failed"
	SecurityActionStateStatusNotStarted SecurityActionStateStatus = "NotStarted"
	SecurityActionStateStatusRunning    SecurityActionStateStatus = "Running"
)

func PossibleValuesForSecurityActionStateStatus() []string {
	return []string{
		string(SecurityActionStateStatusCompleted),
		string(SecurityActionStateStatusFailed),
		string(SecurityActionStateStatusNotStarted),
		string(SecurityActionStateStatusRunning),
	}
}

func parseSecurityActionStateStatus(input string) (*SecurityActionStateStatus, error) {
	vals := map[string]SecurityActionStateStatus{
		"completed":  SecurityActionStateStatusCompleted,
		"failed":     SecurityActionStateStatusFailed,
		"notstarted": SecurityActionStateStatusNotStarted,
		"running":    SecurityActionStateStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityActionStateStatus(input)
	return &out, nil
}
