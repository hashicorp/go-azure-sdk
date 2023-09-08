package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityProfileOperationPredicate struct {
	AzureSubscriptionId  *string
	AzureTenantId        *string
	CreatedDateTime      *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	RiskScore            *string
	UserPrincipalName    *string
}

func (p UserSecurityProfileOperationPredicate) Matches(input UserSecurityProfile) bool {

	if p.AzureSubscriptionId != nil && (input.AzureSubscriptionId == nil || *p.AzureSubscriptionId != *input.AzureSubscriptionId) {
		return false
	}

	if p.AzureTenantId != nil && (input.AzureTenantId == nil || *p.AzureTenantId != *input.AzureTenantId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RiskScore != nil && (input.RiskScore == nil || *p.RiskScore != *input.RiskScore) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
