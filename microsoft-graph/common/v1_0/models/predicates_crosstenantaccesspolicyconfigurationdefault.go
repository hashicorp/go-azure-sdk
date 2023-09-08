package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyConfigurationDefaultOperationPredicate struct {
	Id               *string
	IsServiceDefault *bool
	ODataType        *string
}

func (p CrossTenantAccessPolicyConfigurationDefaultOperationPredicate) Matches(input CrossTenantAccessPolicyConfigurationDefault) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsServiceDefault != nil && (input.IsServiceDefault == nil || *p.IsServiceDefault != *input.IsServiceDefault) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
