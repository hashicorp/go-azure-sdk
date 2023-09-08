package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantIdentitySyncPolicyPartnerOperationPredicate struct {
	DisplayName *string
	ODataType   *string
	TenantId    *string
}

func (p CrossTenantIdentitySyncPolicyPartnerOperationPredicate) Matches(input CrossTenantIdentitySyncPolicyPartner) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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
