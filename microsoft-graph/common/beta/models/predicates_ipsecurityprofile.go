package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpSecurityProfileOperationPredicate struct {
	Address             *string
	AzureSubscriptionId *string
	AzureTenantId       *string
	CountHits           *int64
	CountHosts          *int64
	FirstSeenDateTime   *string
	Id                  *string
	LastSeenDateTime    *string
	ODataType           *string
	RiskScore           *string
}

func (p IpSecurityProfileOperationPredicate) Matches(input IpSecurityProfile) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.AzureSubscriptionId != nil && (input.AzureSubscriptionId == nil || *p.AzureSubscriptionId != *input.AzureSubscriptionId) {
		return false
	}

	if p.AzureTenantId != nil && (input.AzureTenantId == nil || *p.AzureTenantId != *input.AzureTenantId) {
		return false
	}

	if p.CountHits != nil && (input.CountHits == nil || *p.CountHits != *input.CountHits) {
		return false
	}

	if p.CountHosts != nil && (input.CountHosts == nil || *p.CountHosts != *input.CountHosts) {
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

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RiskScore != nil && (input.RiskScore == nil || *p.RiskScore != *input.RiskScore) {
		return false
	}

	return true
}
