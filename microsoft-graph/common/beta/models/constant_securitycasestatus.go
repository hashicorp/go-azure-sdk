package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseStatus string

const (
	SecurityCaseStatusactive          SecurityCaseStatus = "Active"
	SecurityCaseStatusclosed          SecurityCaseStatus = "Closed"
	SecurityCaseStatusclosedWithError SecurityCaseStatus = "ClosedWithError"
	SecurityCaseStatusclosing         SecurityCaseStatus = "Closing"
	SecurityCaseStatuspendingDelete   SecurityCaseStatus = "PendingDelete"
	SecurityCaseStatusunknown         SecurityCaseStatus = "Unknown"
)

func PossibleValuesForSecurityCaseStatus() []string {
	return []string{
		string(SecurityCaseStatusactive),
		string(SecurityCaseStatusclosed),
		string(SecurityCaseStatusclosedWithError),
		string(SecurityCaseStatusclosing),
		string(SecurityCaseStatuspendingDelete),
		string(SecurityCaseStatusunknown),
	}
}

func parseSecurityCaseStatus(input string) (*SecurityCaseStatus, error) {
	vals := map[string]SecurityCaseStatus{
		"active":          SecurityCaseStatusactive,
		"closed":          SecurityCaseStatusclosed,
		"closedwitherror": SecurityCaseStatusclosedWithError,
		"closing":         SecurityCaseStatusclosing,
		"pendingdelete":   SecurityCaseStatuspendingDelete,
		"unknown":         SecurityCaseStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCaseStatus(input)
	return &out, nil
}
