package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobilityManagementPolicyOperationPredicate struct {
	ComplianceUrl *string
	Description   *string
	DiscoveryUrl  *string
	DisplayName   *string
	Id            *string
	IsValid       *bool
	ODataType     *string
	TermsOfUseUrl *string
}

func (p MobilityManagementPolicyOperationPredicate) Matches(input MobilityManagementPolicy) bool {

	if p.ComplianceUrl != nil && (input.ComplianceUrl == nil || *p.ComplianceUrl != *input.ComplianceUrl) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DiscoveryUrl != nil && (input.DiscoveryUrl == nil || *p.DiscoveryUrl != *input.DiscoveryUrl) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsValid != nil && (input.IsValid == nil || *p.IsValid != *input.IsValid) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TermsOfUseUrl != nil && (input.TermsOfUseUrl == nil || *p.TermsOfUseUrl != *input.TermsOfUseUrl) {
		return false
	}

	return true
}
