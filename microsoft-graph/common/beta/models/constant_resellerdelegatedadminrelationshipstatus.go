package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResellerDelegatedAdminRelationshipStatus string

const (
	ResellerDelegatedAdminRelationshipStatusactivating           ResellerDelegatedAdminRelationshipStatus = "Activating"
	ResellerDelegatedAdminRelationshipStatusactive               ResellerDelegatedAdminRelationshipStatus = "Active"
	ResellerDelegatedAdminRelationshipStatusapprovalPending      ResellerDelegatedAdminRelationshipStatus = "ApprovalPending"
	ResellerDelegatedAdminRelationshipStatusapproved             ResellerDelegatedAdminRelationshipStatus = "Approved"
	ResellerDelegatedAdminRelationshipStatuscreated              ResellerDelegatedAdminRelationshipStatus = "Created"
	ResellerDelegatedAdminRelationshipStatusexpired              ResellerDelegatedAdminRelationshipStatus = "Expired"
	ResellerDelegatedAdminRelationshipStatusexpiring             ResellerDelegatedAdminRelationshipStatus = "Expiring"
	ResellerDelegatedAdminRelationshipStatusterminated           ResellerDelegatedAdminRelationshipStatus = "Terminated"
	ResellerDelegatedAdminRelationshipStatusterminating          ResellerDelegatedAdminRelationshipStatus = "Terminating"
	ResellerDelegatedAdminRelationshipStatusterminationRequested ResellerDelegatedAdminRelationshipStatus = "TerminationRequested"
)

func PossibleValuesForResellerDelegatedAdminRelationshipStatus() []string {
	return []string{
		string(ResellerDelegatedAdminRelationshipStatusactivating),
		string(ResellerDelegatedAdminRelationshipStatusactive),
		string(ResellerDelegatedAdminRelationshipStatusapprovalPending),
		string(ResellerDelegatedAdminRelationshipStatusapproved),
		string(ResellerDelegatedAdminRelationshipStatuscreated),
		string(ResellerDelegatedAdminRelationshipStatusexpired),
		string(ResellerDelegatedAdminRelationshipStatusexpiring),
		string(ResellerDelegatedAdminRelationshipStatusterminated),
		string(ResellerDelegatedAdminRelationshipStatusterminating),
		string(ResellerDelegatedAdminRelationshipStatusterminationRequested),
	}
}

func parseResellerDelegatedAdminRelationshipStatus(input string) (*ResellerDelegatedAdminRelationshipStatus, error) {
	vals := map[string]ResellerDelegatedAdminRelationshipStatus{
		"activating":           ResellerDelegatedAdminRelationshipStatusactivating,
		"active":               ResellerDelegatedAdminRelationshipStatusactive,
		"approvalpending":      ResellerDelegatedAdminRelationshipStatusapprovalPending,
		"approved":             ResellerDelegatedAdminRelationshipStatusapproved,
		"created":              ResellerDelegatedAdminRelationshipStatuscreated,
		"expired":              ResellerDelegatedAdminRelationshipStatusexpired,
		"expiring":             ResellerDelegatedAdminRelationshipStatusexpiring,
		"terminated":           ResellerDelegatedAdminRelationshipStatusterminated,
		"terminating":          ResellerDelegatedAdminRelationshipStatusterminating,
		"terminationrequested": ResellerDelegatedAdminRelationshipStatusterminationRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResellerDelegatedAdminRelationshipStatus(input)
	return &out, nil
}
