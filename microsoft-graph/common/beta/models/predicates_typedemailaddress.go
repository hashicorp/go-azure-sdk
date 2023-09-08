package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TypedEmailAddressOperationPredicate struct {
	Address    *string
	Name       *string
	ODataType  *string
	OtherLabel *string
}

func (p TypedEmailAddressOperationPredicate) Matches(input TypedEmailAddress) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OtherLabel != nil && (input.OtherLabel == nil || *p.OtherLabel != *input.OtherLabel) {
		return false
	}

	return true
}
