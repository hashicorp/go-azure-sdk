package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SingleValueLegacyExtendedPropertyOperationPredicate struct {
	Id        *string
	ODataType *string
	Value     *string
}

func (p SingleValueLegacyExtendedPropertyOperationPredicate) Matches(input SingleValueLegacyExtendedProperty) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
