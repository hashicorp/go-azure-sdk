package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipRequestAction string

const (
	DelegatedAdminRelationshipRequestActionapprove         DelegatedAdminRelationshipRequestAction = "Approve"
	DelegatedAdminRelationshipRequestActionlockForApproval DelegatedAdminRelationshipRequestAction = "LockForApproval"
	DelegatedAdminRelationshipRequestActionterminate       DelegatedAdminRelationshipRequestAction = "Terminate"
)

func PossibleValuesForDelegatedAdminRelationshipRequestAction() []string {
	return []string{
		string(DelegatedAdminRelationshipRequestActionapprove),
		string(DelegatedAdminRelationshipRequestActionlockForApproval),
		string(DelegatedAdminRelationshipRequestActionterminate),
	}
}

func parseDelegatedAdminRelationshipRequestAction(input string) (*DelegatedAdminRelationshipRequestAction, error) {
	vals := map[string]DelegatedAdminRelationshipRequestAction{
		"approve":         DelegatedAdminRelationshipRequestActionapprove,
		"lockforapproval": DelegatedAdminRelationshipRequestActionlockForApproval,
		"terminate":       DelegatedAdminRelationshipRequestActionterminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipRequestAction(input)
	return &out, nil
}
