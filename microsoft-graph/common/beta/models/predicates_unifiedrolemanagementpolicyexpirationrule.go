package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyExpirationRuleOperationPredicate struct {
	Id                   *string
	IsExpirationRequired *bool
	MaximumDuration      *string
	ODataType            *string
}

func (p UnifiedRoleManagementPolicyExpirationRuleOperationPredicate) Matches(input UnifiedRoleManagementPolicyExpirationRule) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsExpirationRequired != nil && (input.IsExpirationRequired == nil || *p.IsExpirationRequired != *input.IsExpirationRequired) {
		return false
	}

	if p.MaximumDuration != nil && (input.MaximumDuration == nil || *p.MaximumDuration != *input.MaximumDuration) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
