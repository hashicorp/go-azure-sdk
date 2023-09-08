package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipStatus string

const (
	DelegatedAdminRelationshipStatusactivating           DelegatedAdminRelationshipStatus = "Activating"
	DelegatedAdminRelationshipStatusactive               DelegatedAdminRelationshipStatus = "Active"
	DelegatedAdminRelationshipStatusapprovalPending      DelegatedAdminRelationshipStatus = "ApprovalPending"
	DelegatedAdminRelationshipStatusapproved             DelegatedAdminRelationshipStatus = "Approved"
	DelegatedAdminRelationshipStatuscreated              DelegatedAdminRelationshipStatus = "Created"
	DelegatedAdminRelationshipStatusexpired              DelegatedAdminRelationshipStatus = "Expired"
	DelegatedAdminRelationshipStatusexpiring             DelegatedAdminRelationshipStatus = "Expiring"
	DelegatedAdminRelationshipStatusterminated           DelegatedAdminRelationshipStatus = "Terminated"
	DelegatedAdminRelationshipStatusterminating          DelegatedAdminRelationshipStatus = "Terminating"
	DelegatedAdminRelationshipStatusterminationRequested DelegatedAdminRelationshipStatus = "TerminationRequested"
)

func PossibleValuesForDelegatedAdminRelationshipStatus() []string {
	return []string{
		string(DelegatedAdminRelationshipStatusactivating),
		string(DelegatedAdminRelationshipStatusactive),
		string(DelegatedAdminRelationshipStatusapprovalPending),
		string(DelegatedAdminRelationshipStatusapproved),
		string(DelegatedAdminRelationshipStatuscreated),
		string(DelegatedAdminRelationshipStatusexpired),
		string(DelegatedAdminRelationshipStatusexpiring),
		string(DelegatedAdminRelationshipStatusterminated),
		string(DelegatedAdminRelationshipStatusterminating),
		string(DelegatedAdminRelationshipStatusterminationRequested),
	}
}

func parseDelegatedAdminRelationshipStatus(input string) (*DelegatedAdminRelationshipStatus, error) {
	vals := map[string]DelegatedAdminRelationshipStatus{
		"activating":           DelegatedAdminRelationshipStatusactivating,
		"active":               DelegatedAdminRelationshipStatusactive,
		"approvalpending":      DelegatedAdminRelationshipStatusapprovalPending,
		"approved":             DelegatedAdminRelationshipStatusapproved,
		"created":              DelegatedAdminRelationshipStatuscreated,
		"expired":              DelegatedAdminRelationshipStatusexpired,
		"expiring":             DelegatedAdminRelationshipStatusexpiring,
		"terminated":           DelegatedAdminRelationshipStatusterminated,
		"terminating":          DelegatedAdminRelationshipStatusterminating,
		"terminationrequested": DelegatedAdminRelationshipStatusterminationRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipStatus(input)
	return &out, nil
}
