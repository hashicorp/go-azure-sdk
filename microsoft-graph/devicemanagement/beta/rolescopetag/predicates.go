package rolescopetag

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type RoleScopeTagOperationPredicate struct {
}

func (p RoleScopeTagOperationPredicate) Matches(input beta.RoleScopeTag) bool {

	return true
}

type RoleScopeTagAutoAssignmentOperationPredicate struct {
	Id        *string
	ODataType *string
}

func (p RoleScopeTagAutoAssignmentOperationPredicate) Matches(input beta.RoleScopeTagAutoAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
