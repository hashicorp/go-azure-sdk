package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTimeBasedAttributeTriggerOperationPredicate struct {
	ODataType    *string
	OffsetInDays *int64
}

func (p IdentityGovernanceTimeBasedAttributeTriggerOperationPredicate) Matches(input IdentityGovernanceTimeBasedAttributeTrigger) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OffsetInDays != nil && (input.OffsetInDays == nil || *p.OffsetInDays != *input.OffsetInDays) {
		return false
	}

	return true
}
