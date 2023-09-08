package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanySubscriptionOperationPredicate struct {
	CommerceSubscriptionId *string
	CreatedDateTime        *string
	Id                     *string
	IsTrial                *bool
	NextLifecycleDateTime  *string
	ODataType              *string
	OcpSubscriptionId      *string
	OwnerId                *string
	OwnerTenantId          *string
	OwnerType              *string
	SkuId                  *string
	SkuPartNumber          *string
	Status                 *string
	TotalLicenses          *int64
}

func (p CompanySubscriptionOperationPredicate) Matches(input CompanySubscription) bool {

	if p.CommerceSubscriptionId != nil && (input.CommerceSubscriptionId == nil || *p.CommerceSubscriptionId != *input.CommerceSubscriptionId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsTrial != nil && (input.IsTrial == nil || *p.IsTrial != *input.IsTrial) {
		return false
	}

	if p.NextLifecycleDateTime != nil && (input.NextLifecycleDateTime == nil || *p.NextLifecycleDateTime != *input.NextLifecycleDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OcpSubscriptionId != nil && (input.OcpSubscriptionId == nil || *p.OcpSubscriptionId != *input.OcpSubscriptionId) {
		return false
	}

	if p.OwnerId != nil && (input.OwnerId == nil || *p.OwnerId != *input.OwnerId) {
		return false
	}

	if p.OwnerTenantId != nil && (input.OwnerTenantId == nil || *p.OwnerTenantId != *input.OwnerTenantId) {
		return false
	}

	if p.OwnerType != nil && (input.OwnerType == nil || *p.OwnerType != *input.OwnerType) {
		return false
	}

	if p.SkuId != nil && (input.SkuId == nil || *p.SkuId != *input.SkuId) {
		return false
	}

	if p.SkuPartNumber != nil && (input.SkuPartNumber == nil || *p.SkuPartNumber != *input.SkuPartNumber) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	if p.TotalLicenses != nil && (input.TotalLicenses == nil || *p.TotalLicenses != *input.TotalLicenses) {
		return false
	}

	return true
}
