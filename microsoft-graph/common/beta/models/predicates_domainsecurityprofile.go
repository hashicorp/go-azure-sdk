package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainSecurityProfileOperationPredicate struct {
	AzureSubscriptionId      *string
	AzureTenantId            *string
	CountHits                *int64
	CountInOrg               *int64
	DomainRegisteredDateTime *string
	FirstSeenDateTime        *string
	Id                       *string
	LastSeenDateTime         *string
	Name                     *string
	ODataType                *string
	RiskScore                *string
}

func (p DomainSecurityProfileOperationPredicate) Matches(input DomainSecurityProfile) bool {

	if p.AzureSubscriptionId != nil && (input.AzureSubscriptionId == nil || *p.AzureSubscriptionId != *input.AzureSubscriptionId) {
		return false
	}

	if p.AzureTenantId != nil && (input.AzureTenantId == nil || *p.AzureTenantId != *input.AzureTenantId) {
		return false
	}

	if p.CountHits != nil && (input.CountHits == nil || *p.CountHits != *input.CountHits) {
		return false
	}

	if p.CountInOrg != nil && (input.CountInOrg == nil || *p.CountInOrg != *input.CountInOrg) {
		return false
	}

	if p.DomainRegisteredDateTime != nil && (input.DomainRegisteredDateTime == nil || *p.DomainRegisteredDateTime != *input.DomainRegisteredDateTime) {
		return false
	}

	if p.FirstSeenDateTime != nil && (input.FirstSeenDateTime == nil || *p.FirstSeenDateTime != *input.FirstSeenDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSeenDateTime != nil && (input.LastSeenDateTime == nil || *p.LastSeenDateTime != *input.LastSeenDateTime) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RiskScore != nil && (input.RiskScore == nil || *p.RiskScore != *input.RiskScore) {
		return false
	}

	return true
}
