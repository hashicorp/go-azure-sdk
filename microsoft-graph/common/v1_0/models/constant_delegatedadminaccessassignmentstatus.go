package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminAccessAssignmentStatus string

const (
	DelegatedAdminAccessAssignmentStatusactive   DelegatedAdminAccessAssignmentStatus = "Active"
	DelegatedAdminAccessAssignmentStatusdeleted  DelegatedAdminAccessAssignmentStatus = "Deleted"
	DelegatedAdminAccessAssignmentStatusdeleting DelegatedAdminAccessAssignmentStatus = "Deleting"
	DelegatedAdminAccessAssignmentStatuserror    DelegatedAdminAccessAssignmentStatus = "Error"
	DelegatedAdminAccessAssignmentStatuspending  DelegatedAdminAccessAssignmentStatus = "Pending"
)

func PossibleValuesForDelegatedAdminAccessAssignmentStatus() []string {
	return []string{
		string(DelegatedAdminAccessAssignmentStatusactive),
		string(DelegatedAdminAccessAssignmentStatusdeleted),
		string(DelegatedAdminAccessAssignmentStatusdeleting),
		string(DelegatedAdminAccessAssignmentStatuserror),
		string(DelegatedAdminAccessAssignmentStatuspending),
	}
}

func parseDelegatedAdminAccessAssignmentStatus(input string) (*DelegatedAdminAccessAssignmentStatus, error) {
	vals := map[string]DelegatedAdminAccessAssignmentStatus{
		"active":   DelegatedAdminAccessAssignmentStatusactive,
		"deleted":  DelegatedAdminAccessAssignmentStatusdeleted,
		"deleting": DelegatedAdminAccessAssignmentStatusdeleting,
		"error":    DelegatedAdminAccessAssignmentStatuserror,
		"pending":  DelegatedAdminAccessAssignmentStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminAccessAssignmentStatus(input)
	return &out, nil
}
