package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscribedSkuOperationPredicate struct {
	AccountId        *string
	AccountName      *string
	AppliesTo        *string
	CapabilityStatus *string
	ConsumedUnits    *int64
	Id               *string
	ODataType        *string
	SkuId            *string
	SkuPartNumber    *string
}

func (p SubscribedSkuOperationPredicate) Matches(input SubscribedSku) bool {

	if p.AccountId != nil && (input.AccountId == nil || *p.AccountId != *input.AccountId) {
		return false
	}

	if p.AccountName != nil && (input.AccountName == nil || *p.AccountName != *input.AccountName) {
		return false
	}

	if p.AppliesTo != nil && (input.AppliesTo == nil || *p.AppliesTo != *input.AppliesTo) {
		return false
	}

	if p.CapabilityStatus != nil && (input.CapabilityStatus == nil || *p.CapabilityStatus != *input.CapabilityStatus) {
		return false
	}

	if p.ConsumedUnits != nil && (input.ConsumedUnits == nil || *p.ConsumedUnits != *input.ConsumedUnits) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SkuId != nil && (input.SkuId == nil || *p.SkuId != *input.SkuId) {
		return false
	}

	if p.SkuPartNumber != nil && (input.SkuPartNumber == nil || *p.SkuPartNumber != *input.SkuPartNumber) {
		return false
	}

	return true
}
