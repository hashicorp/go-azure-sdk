package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedOrganizationMembersOperationPredicate struct {
	ConnectedOrganizationId *string
	Description             *string
	ODataType               *string
}

func (p ConnectedOrganizationMembersOperationPredicate) Matches(input ConnectedOrganizationMembers) bool {

	if p.ConnectedOrganizationId != nil && (input.ConnectedOrganizationId == nil || *p.ConnectedOrganizationId != *input.ConnectedOrganizationId) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
