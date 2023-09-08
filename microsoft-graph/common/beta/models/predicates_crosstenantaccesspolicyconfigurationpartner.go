package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyConfigurationPartnerOperationPredicate struct {
	IsInMultiTenantOrganization *bool
	IsServiceProvider           *bool
	ODataType                   *string
	TenantId                    *string
}

func (p CrossTenantAccessPolicyConfigurationPartnerOperationPredicate) Matches(input CrossTenantAccessPolicyConfigurationPartner) bool {

	if p.IsInMultiTenantOrganization != nil && (input.IsInMultiTenantOrganization == nil || *p.IsInMultiTenantOrganization != *input.IsInMultiTenantOrganization) {
		return false
	}

	if p.IsServiceProvider != nil && (input.IsServiceProvider == nil || *p.IsServiceProvider != *input.IsServiceProvider) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
