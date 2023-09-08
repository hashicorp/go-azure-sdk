package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeRuleMembersOperationPredicate struct {
	Description    *string
	MembershipRule *string
	ODataType      *string
}

func (p AttributeRuleMembersOperationPredicate) Matches(input AttributeRuleMembers) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.MembershipRule != nil && (input.MembershipRule == nil || *p.MembershipRule != *input.MembershipRule) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
