package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipOperationStatus string

const (
	DelegatedAdminRelationshipOperationStatusfailed     DelegatedAdminRelationshipOperationStatus = "Failed"
	DelegatedAdminRelationshipOperationStatusnotStarted DelegatedAdminRelationshipOperationStatus = "NotStarted"
	DelegatedAdminRelationshipOperationStatusrunning    DelegatedAdminRelationshipOperationStatus = "Running"
	DelegatedAdminRelationshipOperationStatussucceeded  DelegatedAdminRelationshipOperationStatus = "Succeeded"
)

func PossibleValuesForDelegatedAdminRelationshipOperationStatus() []string {
	return []string{
		string(DelegatedAdminRelationshipOperationStatusfailed),
		string(DelegatedAdminRelationshipOperationStatusnotStarted),
		string(DelegatedAdminRelationshipOperationStatusrunning),
		string(DelegatedAdminRelationshipOperationStatussucceeded),
	}
}

func parseDelegatedAdminRelationshipOperationStatus(input string) (*DelegatedAdminRelationshipOperationStatus, error) {
	vals := map[string]DelegatedAdminRelationshipOperationStatus{
		"failed":     DelegatedAdminRelationshipOperationStatusfailed,
		"notstarted": DelegatedAdminRelationshipOperationStatusnotStarted,
		"running":    DelegatedAdminRelationshipOperationStatusrunning,
		"succeeded":  DelegatedAdminRelationshipOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipOperationStatus(input)
	return &out, nil
}
