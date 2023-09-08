package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedApprovalOperationPredicate struct {
	ApprovalDuration *string
	ApprovalType     *string
	ApproverReason   *string
	EndDateTime      *string
	Id               *string
	ODataType        *string
	RequestorReason  *string
	RoleId           *string
	StartDateTime    *string
	UserId           *string
}

func (p PrivilegedApprovalOperationPredicate) Matches(input PrivilegedApproval) bool {

	if p.ApprovalDuration != nil && (input.ApprovalDuration == nil || *p.ApprovalDuration != *input.ApprovalDuration) {
		return false
	}

	if p.ApprovalType != nil && (input.ApprovalType == nil || *p.ApprovalType != *input.ApprovalType) {
		return false
	}

	if p.ApproverReason != nil && (input.ApproverReason == nil || *p.ApproverReason != *input.ApproverReason) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestorReason != nil && (input.RequestorReason == nil || *p.RequestorReason != *input.RequestorReason) {
		return false
	}

	if p.RoleId != nil && (input.RoleId == nil || *p.RoleId != *input.RoleId) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
