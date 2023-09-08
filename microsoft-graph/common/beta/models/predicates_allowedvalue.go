package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedValueOperationPredicate struct {
	Id        *string
	IsActive  *bool
	ODataType *string
}

func (p AllowedValueOperationPredicate) Matches(input AllowedValue) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsActive != nil && (input.IsActive == nil || *p.IsActive != *input.IsActive) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
