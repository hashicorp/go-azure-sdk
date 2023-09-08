package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADRegistrationPolicyOperationPredicate struct {
	IsAdminConfigurable *bool
	ODataType           *string
}

func (p AzureADRegistrationPolicyOperationPredicate) Matches(input AzureADRegistrationPolicy) bool {

	if p.IsAdminConfigurable != nil && (input.IsAdminConfigurable == nil || *p.IsAdminConfigurable != *input.IsAdminConfigurable) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
