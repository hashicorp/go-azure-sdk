package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleScopeTagOperationPredicate struct {
	Description *string
	DisplayName *string
	Id          *string
	IsBuiltIn   *bool
	ODataType   *string
}

func (p RoleScopeTagOperationPredicate) Matches(input RoleScopeTag) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsBuiltIn != nil && (input.IsBuiltIn == nil || *p.IsBuiltIn != *input.IsBuiltIn) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
