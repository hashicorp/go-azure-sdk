package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerInformationOperationPredicate struct {
	CommerceUrl     *string
	CompanyName     *string
	HelpUrl         *string
	ODataType       *string
	PartnerTenantId *string
	SupportUrl      *string
}

func (p PartnerInformationOperationPredicate) Matches(input PartnerInformation) bool {

	if p.CommerceUrl != nil && (input.CommerceUrl == nil || *p.CommerceUrl != *input.CommerceUrl) {
		return false
	}

	if p.CompanyName != nil && (input.CompanyName == nil || *p.CompanyName != *input.CompanyName) {
		return false
	}

	if p.HelpUrl != nil && (input.HelpUrl == nil || *p.HelpUrl != *input.HelpUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PartnerTenantId != nil && (input.PartnerTenantId == nil || *p.PartnerTenantId != *input.PartnerTenantId) {
		return false
	}

	if p.SupportUrl != nil && (input.SupportUrl == nil || *p.SupportUrl != *input.SupportUrl) {
		return false
	}

	return true
}
