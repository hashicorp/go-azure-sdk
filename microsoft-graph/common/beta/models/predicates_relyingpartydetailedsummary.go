package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelyingPartyDetailedSummaryOperationPredicate struct {
	FailedSignInCount     *int64
	Id                    *string
	ODataType             *string
	RelyingPartyId        *string
	RelyingPartyName      *string
	ServiceId             *string
	SuccessfulSignInCount *int64
	TotalSignInCount      *int64
	UniqueUserCount       *int64
}

func (p RelyingPartyDetailedSummaryOperationPredicate) Matches(input RelyingPartyDetailedSummary) bool {

	if p.FailedSignInCount != nil && (input.FailedSignInCount == nil || *p.FailedSignInCount != *input.FailedSignInCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RelyingPartyId != nil && (input.RelyingPartyId == nil || *p.RelyingPartyId != *input.RelyingPartyId) {
		return false
	}

	if p.RelyingPartyName != nil && (input.RelyingPartyName == nil || *p.RelyingPartyName != *input.RelyingPartyName) {
		return false
	}

	if p.ServiceId != nil && (input.ServiceId == nil || *p.ServiceId != *input.ServiceId) {
		return false
	}

	if p.SuccessfulSignInCount != nil && (input.SuccessfulSignInCount == nil || *p.SuccessfulSignInCount != *input.SuccessfulSignInCount) {
		return false
	}

	if p.TotalSignInCount != nil && (input.TotalSignInCount == nil || *p.TotalSignInCount != *input.TotalSignInCount) {
		return false
	}

	if p.UniqueUserCount != nil && (input.UniqueUserCount == nil || *p.UniqueUserCount != *input.UniqueUserCount) {
		return false
	}

	return true
}
