package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadActionOperationPredicate struct {
	ActionId    *string
	Description *string
	DisplayName *string
	ODataType   *string
	Service     *string
}

func (p ManagedTenantsWorkloadActionOperationPredicate) Matches(input ManagedTenantsWorkloadAction) bool {

	if p.ActionId != nil && (input.ActionId == nil || *p.ActionId != *input.ActionId) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Service != nil && (input.Service == nil || *p.Service != *input.Service) {
		return false
	}

	return true
}
