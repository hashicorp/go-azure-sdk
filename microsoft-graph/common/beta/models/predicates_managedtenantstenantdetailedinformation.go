package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantDetailedInformationOperationPredicate struct {
	City              *string
	CountryCode       *string
	CountryName       *string
	DefaultDomainName *string
	DisplayName       *string
	Id                *string
	IndustryName      *string
	ODataType         *string
	Region            *string
	SegmentName       *string
	TenantId          *string
	VerticalName      *string
}

func (p ManagedTenantsTenantDetailedInformationOperationPredicate) Matches(input ManagedTenantsTenantDetailedInformation) bool {

	if p.City != nil && (input.City == nil || *p.City != *input.City) {
		return false
	}

	if p.CountryCode != nil && (input.CountryCode == nil || *p.CountryCode != *input.CountryCode) {
		return false
	}

	if p.CountryName != nil && (input.CountryName == nil || *p.CountryName != *input.CountryName) {
		return false
	}

	if p.DefaultDomainName != nil && (input.DefaultDomainName == nil || *p.DefaultDomainName != *input.DefaultDomainName) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IndustryName != nil && (input.IndustryName == nil || *p.IndustryName != *input.IndustryName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Region != nil && (input.Region == nil || *p.Region != *input.Region) {
		return false
	}

	if p.SegmentName != nil && (input.SegmentName == nil || *p.SegmentName != *input.SegmentName) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.VerticalName != nil && (input.VerticalName == nil || *p.VerticalName != *input.VerticalName) {
		return false
	}

	return true
}
