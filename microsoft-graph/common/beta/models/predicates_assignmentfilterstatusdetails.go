package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterStatusDetailsOperationPredicate struct {
	ManagedDeviceId *string
	ODataType       *string
	PayloadId       *string
	UserId          *string
}

func (p AssignmentFilterStatusDetailsOperationPredicate) Matches(input AssignmentFilterStatusDetails) bool {

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PayloadId != nil && (input.PayloadId == nil || *p.PayloadId != *input.PayloadId) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
