package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemporaryAccessPassAuthenticationMethodOperationPredicate struct {
	CreatedDateTime       *string
	Id                    *string
	IsUsable              *bool
	IsUsableOnce          *bool
	LifetimeInMinutes     *int64
	MethodUsabilityReason *string
	ODataType             *string
	StartDateTime         *string
	TemporaryAccessPass   *string
}

func (p TemporaryAccessPassAuthenticationMethodOperationPredicate) Matches(input TemporaryAccessPassAuthenticationMethod) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsUsable != nil && (input.IsUsable == nil || *p.IsUsable != *input.IsUsable) {
		return false
	}

	if p.IsUsableOnce != nil && (input.IsUsableOnce == nil || *p.IsUsableOnce != *input.IsUsableOnce) {
		return false
	}

	if p.LifetimeInMinutes != nil && (input.LifetimeInMinutes == nil || *p.LifetimeInMinutes != *input.LifetimeInMinutes) {
		return false
	}

	if p.MethodUsabilityReason != nil && (input.MethodUsabilityReason == nil || *p.MethodUsabilityReason != *input.MethodUsabilityReason) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.TemporaryAccessPass != nil && (input.TemporaryAccessPass == nil || *p.TemporaryAccessPass != *input.TemporaryAccessPass) {
		return false
	}

	return true
}
