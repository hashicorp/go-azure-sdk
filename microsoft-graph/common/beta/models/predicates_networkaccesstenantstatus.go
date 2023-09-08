package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTenantStatusOperationPredicate struct {
	Id                     *string
	ODataType              *string
	OnboardingErrorMessage *string
}

func (p NetworkaccessTenantStatusOperationPredicate) Matches(input NetworkaccessTenantStatus) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnboardingErrorMessage != nil && (input.OnboardingErrorMessage == nil || *p.OnboardingErrorMessage != *input.OnboardingErrorMessage) {
		return false
	}

	return true
}
