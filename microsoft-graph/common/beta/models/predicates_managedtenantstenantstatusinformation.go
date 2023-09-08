package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformationOperationPredicate struct {
	LastDelegatedPrivilegeRefreshDateTime *string
	ODataType                             *string
	OffboardedByUserId                    *string
	OffboardedDateTime                    *string
	OnboardedByUserId                     *string
	OnboardedDateTime                     *string
}

func (p ManagedTenantsTenantStatusInformationOperationPredicate) Matches(input ManagedTenantsTenantStatusInformation) bool {

	if p.LastDelegatedPrivilegeRefreshDateTime != nil && (input.LastDelegatedPrivilegeRefreshDateTime == nil || *p.LastDelegatedPrivilegeRefreshDateTime != *input.LastDelegatedPrivilegeRefreshDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OffboardedByUserId != nil && (input.OffboardedByUserId == nil || *p.OffboardedByUserId != *input.OffboardedByUserId) {
		return false
	}

	if p.OffboardedDateTime != nil && (input.OffboardedDateTime == nil || *p.OffboardedDateTime != *input.OffboardedDateTime) {
		return false
	}

	if p.OnboardedByUserId != nil && (input.OnboardedByUserId == nil || *p.OnboardedByUserId != *input.OnboardedByUserId) {
		return false
	}

	if p.OnboardedDateTime != nil && (input.OnboardedDateTime == nil || *p.OnboardedDateTime != *input.OnboardedDateTime) {
		return false
	}

	return true
}
