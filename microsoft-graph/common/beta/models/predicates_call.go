package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallOperationPredicate struct {
	CallChainId             *string
	CallbackUri             *string
	Id                      *string
	MyParticipantId         *string
	ODataType               *string
	RingingTimeoutInSeconds *int64
	Subject                 *string
	TenantId                *string
	TerminationReason       *string
}

func (p CallOperationPredicate) Matches(input Call) bool {

	if p.CallChainId != nil && (input.CallChainId == nil || *p.CallChainId != *input.CallChainId) {
		return false
	}

	if p.CallbackUri != nil && (input.CallbackUri == nil || *p.CallbackUri != *input.CallbackUri) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MyParticipantId != nil && (input.MyParticipantId == nil || *p.MyParticipantId != *input.MyParticipantId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RingingTimeoutInSeconds != nil && (input.RingingTimeoutInSeconds == nil || *p.RingingTimeoutInSeconds != *input.RingingTimeoutInSeconds) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.TerminationReason != nil && (input.TerminationReason == nil || *p.TerminationReason != *input.TerminationReason) {
		return false
	}

	return true
}
