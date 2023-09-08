package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenLicenseSummaryOperationPredicate struct {
	AppleId               *string
	AvailableLicenseCount *int64
	ODataType             *string
	OrganizationName      *string
	UsedLicenseCount      *int64
	VppTokenId            *string
}

func (p VppTokenLicenseSummaryOperationPredicate) Matches(input VppTokenLicenseSummary) bool {

	if p.AppleId != nil && (input.AppleId == nil || *p.AppleId != *input.AppleId) {
		return false
	}

	if p.AvailableLicenseCount != nil && (input.AvailableLicenseCount == nil || *p.AvailableLicenseCount != *input.AvailableLicenseCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OrganizationName != nil && (input.OrganizationName == nil || *p.OrganizationName != *input.OrganizationName) {
		return false
	}

	if p.UsedLicenseCount != nil && (input.UsedLicenseCount == nil || *p.UsedLicenseCount != *input.UsedLicenseCount) {
		return false
	}

	if p.VppTokenId != nil && (input.VppTokenId == nil || *p.VppTokenId != *input.VppTokenId) {
		return false
	}

	return true
}
