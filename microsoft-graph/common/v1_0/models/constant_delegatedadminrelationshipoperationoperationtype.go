package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipOperationOperationType string

const (
	DelegatedAdminRelationshipOperationOperationTypedelegatedAdminAccessAssignmentUpdate DelegatedAdminRelationshipOperationOperationType = "DelegatedAdminAccessAssignmentUpdate"
)

func PossibleValuesForDelegatedAdminRelationshipOperationOperationType() []string {
	return []string{
		string(DelegatedAdminRelationshipOperationOperationTypedelegatedAdminAccessAssignmentUpdate),
	}
}

func parseDelegatedAdminRelationshipOperationOperationType(input string) (*DelegatedAdminRelationshipOperationOperationType, error) {
	vals := map[string]DelegatedAdminRelationshipOperationOperationType{
		"delegatedadminaccessassignmentupdate": DelegatedAdminRelationshipOperationOperationTypedelegatedAdminAccessAssignmentUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipOperationOperationType(input)
	return &out, nil
}
