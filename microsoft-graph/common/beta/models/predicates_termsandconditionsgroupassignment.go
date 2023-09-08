package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionsGroupAssignmentOperationPredicate struct {
	Id            *string
	ODataType     *string
	TargetGroupId *string
}

func (p TermsAndConditionsGroupAssignmentOperationPredicate) Matches(input TermsAndConditionsGroupAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TargetGroupId != nil && (input.TargetGroupId == nil || *p.TargetGroupId != *input.TargetGroupId) {
		return false
	}

	return true
}
