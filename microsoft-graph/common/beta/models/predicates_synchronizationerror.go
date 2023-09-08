package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationErrorOperationPredicate struct {
	Code             *string
	Message          *string
	ODataType        *string
	TenantActionable *bool
}

func (p SynchronizationErrorOperationPredicate) Matches(input SynchronizationError) bool {

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantActionable != nil && (input.TenantActionable == nil || *p.TenantActionable != *input.TenantActionable) {
		return false
	}

	return true
}
