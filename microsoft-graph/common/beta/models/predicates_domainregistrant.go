package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainRegistrantOperationPredicate struct {
	CountryOrRegionCode *string
	ODataType           *string
	Organization        *string
	Url                 *string
	Vendor              *string
}

func (p DomainRegistrantOperationPredicate) Matches(input DomainRegistrant) bool {

	if p.CountryOrRegionCode != nil && (input.CountryOrRegionCode == nil || *p.CountryOrRegionCode != *input.CountryOrRegionCode) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Organization != nil && (input.Organization == nil || *p.Organization != *input.Organization) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	if p.Vendor != nil && (input.Vendor == nil || *p.Vendor != *input.Vendor) {
		return false
	}

	return true
}
