package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseStatus string

const (
	EdiscoveryCaseStatusactive          EdiscoveryCaseStatus = "Active"
	EdiscoveryCaseStatusclosed          EdiscoveryCaseStatus = "Closed"
	EdiscoveryCaseStatusclosedWithError EdiscoveryCaseStatus = "ClosedWithError"
	EdiscoveryCaseStatusclosing         EdiscoveryCaseStatus = "Closing"
	EdiscoveryCaseStatuspendingDelete   EdiscoveryCaseStatus = "PendingDelete"
	EdiscoveryCaseStatusunknown         EdiscoveryCaseStatus = "Unknown"
)

func PossibleValuesForEdiscoveryCaseStatus() []string {
	return []string{
		string(EdiscoveryCaseStatusactive),
		string(EdiscoveryCaseStatusclosed),
		string(EdiscoveryCaseStatusclosedWithError),
		string(EdiscoveryCaseStatusclosing),
		string(EdiscoveryCaseStatuspendingDelete),
		string(EdiscoveryCaseStatusunknown),
	}
}

func parseEdiscoveryCaseStatus(input string) (*EdiscoveryCaseStatus, error) {
	vals := map[string]EdiscoveryCaseStatus{
		"active":          EdiscoveryCaseStatusactive,
		"closed":          EdiscoveryCaseStatusclosed,
		"closedwitherror": EdiscoveryCaseStatusclosedWithError,
		"closing":         EdiscoveryCaseStatusclosing,
		"pendingdelete":   EdiscoveryCaseStatuspendingDelete,
		"unknown":         EdiscoveryCaseStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseStatus(input)
	return &out, nil
}
