package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeAssignmentOperationPredicate struct {
	DisplayName          *string
	Id                   *string
	IsOptional           *bool
	ODataType            *string
	RequiresVerification *bool
}

func (p IdentityUserFlowAttributeAssignmentOperationPredicate) Matches(input IdentityUserFlowAttributeAssignment) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsOptional != nil && (input.IsOptional == nil || *p.IsOptional != *input.IsOptional) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequiresVerification != nil && (input.RequiresVerification == nil || *p.RequiresVerification != *input.RequiresVerification) {
		return false
	}

	return true
}
