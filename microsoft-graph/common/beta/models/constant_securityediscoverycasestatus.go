package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCaseStatus string

const (
	SecurityEdiscoveryCaseStatusactive          SecurityEdiscoveryCaseStatus = "Active"
	SecurityEdiscoveryCaseStatusclosed          SecurityEdiscoveryCaseStatus = "Closed"
	SecurityEdiscoveryCaseStatusclosedWithError SecurityEdiscoveryCaseStatus = "ClosedWithError"
	SecurityEdiscoveryCaseStatusclosing         SecurityEdiscoveryCaseStatus = "Closing"
	SecurityEdiscoveryCaseStatuspendingDelete   SecurityEdiscoveryCaseStatus = "PendingDelete"
	SecurityEdiscoveryCaseStatusunknown         SecurityEdiscoveryCaseStatus = "Unknown"
)

func PossibleValuesForSecurityEdiscoveryCaseStatus() []string {
	return []string{
		string(SecurityEdiscoveryCaseStatusactive),
		string(SecurityEdiscoveryCaseStatusclosed),
		string(SecurityEdiscoveryCaseStatusclosedWithError),
		string(SecurityEdiscoveryCaseStatusclosing),
		string(SecurityEdiscoveryCaseStatuspendingDelete),
		string(SecurityEdiscoveryCaseStatusunknown),
	}
}

func parseSecurityEdiscoveryCaseStatus(input string) (*SecurityEdiscoveryCaseStatus, error) {
	vals := map[string]SecurityEdiscoveryCaseStatus{
		"active":          SecurityEdiscoveryCaseStatusactive,
		"closed":          SecurityEdiscoveryCaseStatusclosed,
		"closedwitherror": SecurityEdiscoveryCaseStatusclosedWithError,
		"closing":         SecurityEdiscoveryCaseStatusclosing,
		"pendingdelete":   SecurityEdiscoveryCaseStatuspendingDelete,
		"unknown":         SecurityEdiscoveryCaseStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryCaseStatus(input)
	return &out, nil
}
