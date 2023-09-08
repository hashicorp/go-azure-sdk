package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyOperationOperationStatus string

const (
	GroupPolicyOperationOperationStatusfailed     GroupPolicyOperationOperationStatus = "Failed"
	GroupPolicyOperationOperationStatusinProgress GroupPolicyOperationOperationStatus = "InProgress"
	GroupPolicyOperationOperationStatussuccess    GroupPolicyOperationOperationStatus = "Success"
	GroupPolicyOperationOperationStatusunknown    GroupPolicyOperationOperationStatus = "Unknown"
)

func PossibleValuesForGroupPolicyOperationOperationStatus() []string {
	return []string{
		string(GroupPolicyOperationOperationStatusfailed),
		string(GroupPolicyOperationOperationStatusinProgress),
		string(GroupPolicyOperationOperationStatussuccess),
		string(GroupPolicyOperationOperationStatusunknown),
	}
}

func parseGroupPolicyOperationOperationStatus(input string) (*GroupPolicyOperationOperationStatus, error) {
	vals := map[string]GroupPolicyOperationOperationStatus{
		"failed":     GroupPolicyOperationOperationStatusfailed,
		"inprogress": GroupPolicyOperationOperationStatusinProgress,
		"success":    GroupPolicyOperationOperationStatussuccess,
		"unknown":    GroupPolicyOperationOperationStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyOperationOperationStatus(input)
	return &out, nil
}
