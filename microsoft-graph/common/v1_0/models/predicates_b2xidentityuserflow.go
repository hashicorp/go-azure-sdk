package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xIdentityUserFlowOperationPredicate struct {
	Id        *string
	ODataType *string
}

func (p B2xIdentityUserFlowOperationPredicate) Matches(input B2xIdentityUserFlow) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
