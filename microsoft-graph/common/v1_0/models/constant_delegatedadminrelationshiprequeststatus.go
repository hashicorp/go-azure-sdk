package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipRequestStatus string

const (
	DelegatedAdminRelationshipRequestStatuscreated   DelegatedAdminRelationshipRequestStatus = "Created"
	DelegatedAdminRelationshipRequestStatusfailed    DelegatedAdminRelationshipRequestStatus = "Failed"
	DelegatedAdminRelationshipRequestStatuspending   DelegatedAdminRelationshipRequestStatus = "Pending"
	DelegatedAdminRelationshipRequestStatussucceeded DelegatedAdminRelationshipRequestStatus = "Succeeded"
)

func PossibleValuesForDelegatedAdminRelationshipRequestStatus() []string {
	return []string{
		string(DelegatedAdminRelationshipRequestStatuscreated),
		string(DelegatedAdminRelationshipRequestStatusfailed),
		string(DelegatedAdminRelationshipRequestStatuspending),
		string(DelegatedAdminRelationshipRequestStatussucceeded),
	}
}

func parseDelegatedAdminRelationshipRequestStatus(input string) (*DelegatedAdminRelationshipRequestStatus, error) {
	vals := map[string]DelegatedAdminRelationshipRequestStatus{
		"created":   DelegatedAdminRelationshipRequestStatuscreated,
		"failed":    DelegatedAdminRelationshipRequestStatusfailed,
		"pending":   DelegatedAdminRelationshipRequestStatuspending,
		"succeeded": DelegatedAdminRelationshipRequestStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipRequestStatus(input)
	return &out, nil
}
