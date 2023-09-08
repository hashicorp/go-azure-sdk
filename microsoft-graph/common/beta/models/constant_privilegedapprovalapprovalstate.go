package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedApprovalApprovalState string

const (
	PrivilegedApprovalApprovalStateaborted  PrivilegedApprovalApprovalState = "Aborted"
	PrivilegedApprovalApprovalStateapproved PrivilegedApprovalApprovalState = "Approved"
	PrivilegedApprovalApprovalStatecanceled PrivilegedApprovalApprovalState = "Canceled"
	PrivilegedApprovalApprovalStatedenied   PrivilegedApprovalApprovalState = "Denied"
	PrivilegedApprovalApprovalStatepending  PrivilegedApprovalApprovalState = "Pending"
)

func PossibleValuesForPrivilegedApprovalApprovalState() []string {
	return []string{
		string(PrivilegedApprovalApprovalStateaborted),
		string(PrivilegedApprovalApprovalStateapproved),
		string(PrivilegedApprovalApprovalStatecanceled),
		string(PrivilegedApprovalApprovalStatedenied),
		string(PrivilegedApprovalApprovalStatepending),
	}
}

func parsePrivilegedApprovalApprovalState(input string) (*PrivilegedApprovalApprovalState, error) {
	vals := map[string]PrivilegedApprovalApprovalState{
		"aborted":  PrivilegedApprovalApprovalStateaborted,
		"approved": PrivilegedApprovalApprovalStateapproved,
		"canceled": PrivilegedApprovalApprovalStatecanceled,
		"denied":   PrivilegedApprovalApprovalStatedenied,
		"pending":  PrivilegedApprovalApprovalStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedApprovalApprovalState(input)
	return &out, nil
}
